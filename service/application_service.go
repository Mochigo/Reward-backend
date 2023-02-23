package service

import (
	"Reward/common"
	"Reward/model"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
)

type ApplicationService struct {
	Ctx            *gin.Context
	applicationDao *model.ApplicationDao
}

func NewApplicationService(ctx *gin.Context) *ApplicationService {
	return &ApplicationService{
		Ctx:            ctx,
		applicationDao: model.GetApplicationDao(),
	}
}

func (s *ApplicationService) CreateApplication(req *entity.CreateApplicationEntity) error {
	return s.applicationDao.Create(model.DB.Self, &model.Application{
		ScholarshipItemId: req.ScholarshipItemId,
		ScholarshipId:     req.ScholarshipId,
		StudentId:         s.Ctx.GetInt64("userID"),
		Status:            common.StatusPROCESS,
	})
}
