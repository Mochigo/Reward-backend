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
	"Reward/log"
	"Reward/model"
	"Reward/service/entity"
)

type StudentService struct {
	ctx        *gin.Context
	studentDao *model.StudentDao
	collegeDao *model.CollegeDao
}

func NewStudentService(ctx *gin.Context) *StudentService {
	return &StudentService{
		ctx:        ctx,
		studentDao: model.GetStudentDao(),
		collegeDao: model.GetCollegeDao(),
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
