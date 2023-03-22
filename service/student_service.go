package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/common"
	"Reward/common/token"
	"Reward/common/utils"
	"Reward/log"
	"Reward/model"
	"Reward/service/entity"
)

type StudentService struct {
	ctx                *gin.Context
	studentDao         *model.StudentDao
	collegeDao         *model.CollegeDao
	applicationDao     *model.ApplicationDao
	scholarshipItemDao *model.ScholarshipItemDao
	scholarshipDao     *model.ScholarshipDao
}

func NewStudentService(ctx *gin.Context) *StudentService {
	return &StudentService{
		ctx:                ctx,
		studentDao:         model.GetStudentDao(),
		collegeDao:         model.GetCollegeDao(),
		applicationDao:     model.GetApplicationDao(),
		scholarshipItemDao: model.GetScholarshipItemDao(),
		scholarshipDao:     model.GetScholarshipDao(),
	}
}

// 判断当前uid是否存在，且密码与之对应
func (s *StudentService) VerifyStudent(uid, password string) (bool, error) {
	stu, err := s.studentDao.GetStudentByUID(model.DB.Self, uid)
	if err != nil {
		return false, err
	}

	if stu.Password != password {
		log.Error("[VerifyStudent]the accound does not match the password",
			zap.String("error", common.ErrMismatching.Error()))
		return false, common.ErrMismatching
	}

	return true, nil
}

func (s *StudentService) Sign(uid string) (string, error) {
	stu, err := s.studentDao.GetStudentByUID(model.DB.Self, uid)
	if err != nil {
		return common.StringEmpty, err
	}

	t, err := token.GenerateToken(&token.PayLoad{
		UserID:    int(stu.Id),
		CollegeId: int(stu.CollegeId),
		Expired:   time.Duration(viper.GetInt("token_expired") * int(time.Second)),
	})
	if err != nil {
		return common.StringEmpty, err
	}

	return t, nil
}

func (s *StudentService) GetStudentInfo(sid int64) (*entity.StudentEntity, error) {
	stu, err := s.studentDao.GetStudentById(model.DB.Self, sid)
	if err != nil {
		return nil, err
	}

	college, err := s.collegeDao.GetCollegeById(model.DB.Self, stu.CollegeId)
	if err != nil {
		return nil, err
	}
	return &entity.StudentEntity{
		Id:      stu.Id,
		Score:   stu.Score,
		Uid:     stu.Uid,
		College: college.Name,
	}, nil
}

func (s *StudentService) UploadScore(students []*entity.CreateStudentEntity) (*entity.UploadScoreEntity, error) {
	successfulList := []string{}
	failedList := []string{}
	for _, stu := range students {
		if err := s.SaveStudent(stu); err != nil {
			failedList = append(failedList, stu.Uid)
			continue
		}
		successfulList = append(successfulList, stu.Uid)
	}

	return &entity.UploadScoreEntity{
		SuccessfulList: successfulList,
		FailedList:     failedList,
	}, nil
}

func (s *StudentService) SaveStudent(entity *entity.CreateStudentEntity) error {
	stu, err := s.studentDao.GetStudentByUID(model.DB.Self, entity.Uid)
	// 记录不存在的情况
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s.studentDao.Create(model.DB.Self, &model.Student{
			Score:     entity.Score,
			Uid:       entity.Uid,
			Password:  entity.Password,
			CollegeId: int64(entity.CollegeId),
		})
	}

	stu.Score = entity.Score
	return s.studentDao.Save(model.DB.Self, stu)
}

func (s *StudentService) GetUserApplication(req *entity.GetUserApplicationEntity) ([]*entity.ApplicationEntity, int64, error) {
	condi := make(map[string]interface{})
	condi[common.CondiPage] = req.Page
	condi[common.CondiLimit] = req.Limit
	condi[common.CondiStudentId] = s.ctx.GetInt(common.TokenUserID)
	// 查询申请
	al, err := s.applicationDao.GetList(model.DB.Self, condi)
	if err != nil {
		return nil, 0, err
	}

	scholarship := make(map[int64]*model.Scholarship)
	scholarshipIds := make([]int64, 0, len(al))

	scholarshipItem := make(map[int64]*model.ScholarshipItem)
	scholarshipItemIds := make([]int64, 0, len(al))

	for _, a := range al {
		scholarshipItemIds = append(scholarshipItemIds, a.ScholarshipItemId)
		if _, ok := scholarship[a.ScholarshipId]; !ok {
			scholarshipIds = append(scholarshipIds, a.ScholarshipId)
			scholarship[a.ScholarshipId] = &model.Scholarship{}
		}
	}

	scholarshipItems, err := s.scholarshipItemDao.BatchGetByIds(model.DB.Self, scholarshipItemIds)
	if err != nil {
		return nil, 0, err
	}
	for _, item := range scholarshipItems {
		scholarshipItem[item.Id] = item
	}

	scholarships, err := s.scholarshipDao.BatchGetByIds(model.DB.Self, scholarshipIds)
	if err != nil {
		return nil, 0, err
	}
	for _, s := range scholarships {
		scholarship[s.Id].Name = s.Name
		scholarship[s.Id].CollegeId = s.CollegeId
		scholarship[s.Id].EndTime = s.EndTime
		scholarship[s.Id].StartTime = s.StartTime
	}

	// 返回拼接申请内容
	applications := make([]*entity.ApplicationEntity, 0, len(al))
	for _, a := range al {
		if _, ok := scholarshipItem[a.ScholarshipItemId]; !ok {
			continue
		}
		if _, ok := scholarship[a.ScholarshipId]; !ok {
			continue
		}
		tmp := &entity.ApplicationEntity{
			Id:                  a.Id,
			ScholarshipItemId:   a.ScholarshipItemId,
			ScholarshipItemName: scholarshipItem[a.ScholarshipItemId].Name,
			ScholarshipId:       a.ScholarshipId,
			ScholarshipName:     scholarship[a.ScholarshipId].Name,
			StudentId:           a.StudentId,
			Status:              a.Status,
			Deadline:            a.Deadline.Format(utils.LayoutDateTime),
		}
		applications = append(applications, tmp)
	}

	total, err := s.applicationDao.GetCountByStudentId(model.DB.Self, s.ctx.GetInt(common.TokenUserID))
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}
