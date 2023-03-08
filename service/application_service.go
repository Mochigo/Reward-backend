package service

import (
	"Reward/common"
	"Reward/model"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
)

type ApplicationService struct {
	Ctx                *gin.Context
	applicationDao     *model.ApplicationDao
	scholarshipItemDao *model.ScholarshipItemDao
}

func NewApplicationService(ctx *gin.Context) *ApplicationService {
	return &ApplicationService{
		Ctx:                ctx,
		applicationDao:     model.GetApplicationDao(),
		scholarshipItemDao: model.GetScholarshipItemDao(),
	}
}

func (s *ApplicationService) CreateApplication(req *entity.CreateApplicationEntity) error {
	return s.applicationDao.Create(model.DB.Self, &model.Application{
		ScholarshipItemId: req.ScholarshipItemId,
		ScholarshipId:     req.ScholarshipId,
		StudentId:         int64(s.Ctx.GetInt(common.TokenUserID)),
		Status:            common.StatusPROCESS,
	})
}

func (s *ApplicationService) GetUserApplication(req *entity.GetUserApplicationEntity) ([]*entity.ApplicationEntity, int64, error) {
	condi := make(map[string]interface{})
	condi[common.CondiPage] = req.Page
	condi[common.CondiLimit] = req.Limit
	condi[common.CondiStudentId] = s.Ctx.GetInt(common.TokenUserID)

	al, err := s.applicationDao.GetList(model.DB.Self, condi)
	if err != nil {
		return nil, 0, err
	}

	scholarshipItemIds := make([]int64, 0, len(al))
	for _, a := range al {
		scholarshipItemIds = append(scholarshipItemIds, a.ScholarshipItemId)
	}
	scholarshipItems, err := s.scholarshipItemDao.BatchGetByIds(model.DB.Self, scholarshipItemIds)
	if err != nil {
		return nil, 0, err
	}
	scholarshipItem := make(map[int64]*model.ScholarshipItem)
	for _, item := range scholarshipItems {
		scholarshipItem[item.Id] = item
	}

	applications := make([]*entity.ApplicationEntity, 0, len(al))
	for _, a := range al {
		tmp := &entity.ApplicationEntity{
			Id:                  a.Id,
			ScholarshipItemId:   a.ScholarshipItemId,
			ScholarshipItemName: scholarshipItem[a.ScholarshipItemId].Name,
			ScholarshipId:       a.ScholarshipId,
			StudentId:           a.StudentId,
			Status:              a.Status,
		}
		applications = append(applications, tmp)
	}

	total, err := s.applicationDao.GetCountByStudentId(model.DB.Self, s.Ctx.GetInt(common.TokenUserID))
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}
