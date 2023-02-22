package service

import (
	"Reward/common"
	"Reward/common/utils"
	"Reward/model"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
)

type ScholarshipService struct {
	ctx            *gin.Context
	scholarshipDao *model.ScholarshipDao
	attchmentDao   *model.AttachmentDao
}

func NewScholarshipService(ctx *gin.Context) *ScholarshipService {
	return &ScholarshipService{
		ctx:            ctx,
		scholarshipDao: model.GetScholarshipDao(),
		attchmentDao:   model.GetAttachmentDao(),
	}
}

func (s *ScholarshipService) AddAttachment(req *entity.AddAttachmentEntity) error {
	return s.attchmentDao.Create(model.DB.Self, &model.Attachment{
		ScholarshipId: req.ScholarshipId,
		Url:           req.Url,
	})
}

func (s *ScholarshipService) GetAttachments(req *entity.GetAttachmentsEntity) ([]*entity.AttachmentEntity, error) {
	al, err := s.attchmentDao.GetListByScholarshipId(model.DB.Self, req.ScholarshipId)
	if err != nil {
		return nil, err
	}

	attachments := make([]*entity.AttachmentEntity, 0)
	for _, a := range al {
		tmp := &entity.AttachmentEntity{
			Url: a.Url,
		}
		attachments = append(attachments, tmp)
	}
	return attachments, nil
}

func (s *ScholarshipService) CreateScholarship(req *entity.CreateScholarshipEntity) error {
	start, err := utils.GetDateTime(req.StartTime)
	if err != nil {
		return common.ErrTimeParse
	}

	end, err := utils.GetDateTime(req.EndTime)
	if err != nil {
		return common.ErrTimeParse
	}

	return s.scholarshipDao.Create(model.DB.Self, &model.Scholarship{
		Name:      req.Name,
		CollegeId: s.ctx.GetInt64("collegeID"),
		StartTime: start,
		EndTime:   end,
	})
}
