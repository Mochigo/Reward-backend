package service

import (
	"Reward/common"
	"Reward/common/utils"
	"Reward/log"
	"Reward/model"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApplicationService struct {
	Ctx            *gin.Context
	applicationDao *model.ApplicationDao
	studentDao     *model.StudentDao
}

func NewApplicationService(ctx *gin.Context) *ApplicationService {
	return &ApplicationService{
		Ctx:            ctx,
		applicationDao: model.GetApplicationDao(),
		studentDao:     model.GetStudentDao(),
	}
}

func (s *ApplicationService) CreateApplication(req *entity.CreateApplicationEntity) error {
	deadline, err := utils.GetDateTime(req.Deadline)
	if err != nil {
		log.Error("[CreateApplication] fail to parse time",
			zap.String("time", req.Deadline))
		return common.ErrTimeParse
	}

	return s.applicationDao.Create(model.DB.Self, &model.Application{
		ScholarshipItemId: req.ScholarshipItemId,
		ScholarshipId:     req.ScholarshipId,
		StudentId:         int64(s.Ctx.GetInt(common.TokenUserID)),
		Status:            common.StatusPROCESS,
		Deadline:          deadline,
	})
}

func (s *ApplicationService) GetItemApplications(req *entity.GetItemApplicationsEntity) ([]*entity.ItemApplicationEntity, int64, error) {
	al, err := s.applicationDao.GetListByScholarshipItemId(model.DB.Self, req.ScholarshipItemId, req.Page, req.Limit)
	if err != nil {
		return nil, 0, err
	}

	stuIds := make([]int64, 0, len(al))
	for _, a := range al {
		stuIds = append(stuIds, a.StudentId)
	}

	students, err := s.studentDao.BatchGetByIds(model.DB.Self, stuIds)
	if err != nil {
		return nil, 0, err
	}
	student := make(map[int64]*model.Student)
	for _, s := range students {
		student[s.Id] = s
	}

	// 返回拼接申请内容
	applications := make([]*entity.ItemApplicationEntity, 0, len(al))
	for _, a := range al {
		if _, exist := student[a.StudentId]; !exist {
			continue
		}
		tmp := &entity.ItemApplicationEntity{
			Id:                a.Id,
			ScholarshipItemId: a.ScholarshipItemId,
			ScholarshipId:     a.ScholarshipId,
			StudentId:         a.StudentId,
			Uid:               student[a.StudentId].Uid,
			Status:            a.Status,
			Deadline:          a.Deadline.Format(utils.LayoutDateTime),
		}
		applications = append(applications, tmp)
	}

	total, err := s.applicationDao.GetCountByScholarshipItemId(model.DB.Self, req.ScholarshipItemId)
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

func (s *ApplicationService) AuditApplication(req *entity.AuditApplicationEntity) error {
	updateMaps := make(map[string]interface{})
	updateMaps["status"] = common.StatusAPPROVE

	return s.applicationDao.Update(model.DB.Self, req.ApplicationId, updateMaps)
}
