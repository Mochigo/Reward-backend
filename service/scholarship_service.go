package service

import (
	"github.com/gin-gonic/gin"

	"Reward/common"
	"Reward/common/utils"
	"Reward/model"
	"Reward/service/entity"
)

type ScholarshipService struct {
	ctx                *gin.Context
	scholarshipDao     *model.ScholarshipDao
	attchmentDao       *model.AttachmentDao
	scholarshipItemDao *model.ScholarshipItemDao
	applicationDao     *model.ApplicationDao
}

func NewScholarshipService(ctx *gin.Context) *ScholarshipService {
	return &ScholarshipService{
		ctx:                ctx,
		scholarshipDao:     model.GetScholarshipDao(),
		attchmentDao:       model.GetAttachmentDao(),
		applicationDao:     model.GetApplicationDao(),
		scholarshipItemDao: model.GetScholarshipItemDao(),
	}
}

func (s *ScholarshipService) AddAttachment(req *entity.AddAttachmentEntity) error {
	return s.attchmentDao.Create(model.DB.Self, &model.Attachment{
		ScholarshipId: req.ScholarshipId,
		Url:           req.Url,
	})
}

func (s *ScholarshipService) RemoveAttachment(req *entity.RemoveAttachmentEntity) error {

	return s.attchmentDao.DeleteByID(model.DB.Self, req.AttachmentId)

}

func (s *ScholarshipService) GetAttachments(req *entity.GetAttachmentsEntity) ([]*entity.AttachmentEntity, error) {
	al, err := s.attchmentDao.GetListByScholarshipId(model.DB.Self, req.ScholarshipId)
	if err != nil {
		return nil, err
	}

	attachments := make([]*entity.AttachmentEntity, 0, len(al))
	for _, a := range al {
		tmp := &entity.AttachmentEntity{
			Id:  a.Id,
			Url: a.Url,
		}
		attachments = append(attachments, tmp)
	}
	return attachments, nil
}

func (s *ScholarshipService) CreateScholarship(req *entity.CreateScholarshipEntity) (int64, error) {
	start, err := utils.GetDateTime(req.StartTime)
	if err != nil {
		return 0, common.ErrTimeParse
	}

	end, err := utils.GetDateTime(req.EndTime)
	if err != nil {
		return 0, common.ErrTimeParse
	}

	return s.scholarshipDao.Create(model.DB.Self, &model.Scholarship{
		Name:      req.Name,
		CollegeId: int64(s.ctx.GetInt(common.TokenCollegeID)),
		StartTime: start,
		EndTime:   end,
	})
}

func (s *ScholarshipService) GetScholarshipInfo(req *entity.GetScholarshipInfoEntity) (*entity.ScholarshipEntity, error) {
	scholarship, err := s.scholarshipDao.GetScholarshipById(model.DB.Self, req.ScholarshipId)
	if err != nil {
		return nil, err
	}

	return &entity.ScholarshipEntity{
		Id:        scholarship.Id,
		Name:      scholarship.Name,
		CollegeId: scholarship.CollegeId,
		StartTime: scholarship.StartTime.Format(utils.LayoutDateTime),
		EndTime:   scholarship.EndTime.Format(utils.LayoutDateTime),
	}, nil
}

func (s *ScholarshipService) GetScholarships(req *entity.GetScholarshipsEntity) ([]*entity.ScholarshipEntity, int64, error) {
	condi := make(map[string]interface{})
	condi[common.CondiPage] = req.Page
	condi[common.CondiLimit] = req.Limit
	condi[common.CondiCollegeId] = s.ctx.GetInt(common.TokenCollegeID)

	sl, err := s.scholarshipDao.GetList(model.DB.Self, condi)
	if err != nil {
		return nil, 0, err
	}

	scholarships := make([]*entity.ScholarshipEntity, 0, len(sl))
	for _, s := range sl {
		tmp := &entity.ScholarshipEntity{
			Id:        s.Id,
			Name:      s.Name,
			CollegeId: s.CollegeId,
			StartTime: s.StartTime.Format(utils.LayoutDateTime),
			EndTime:   s.EndTime.Format(utils.LayoutDateTime),
		}

		scholarships = append(scholarships, tmp)
	}

	total, err := s.scholarshipDao.GetCountByCollegeId(model.DB.Self, s.ctx.GetInt(common.TokenCollegeID))
	if err != nil {
		return nil, 0, err
	}
	return scholarships, total, nil
}

func (s *ScholarshipService) AddScholarshipItem(req *entity.AddScholarshipItemEntity) error {
	return s.scholarshipItemDao.Create(model.DB.Self, &model.ScholarshipItem{
		Name:          req.Name,
		ScholarshipId: req.ScholarshipId,
	})
}

func (s *ScholarshipService) GetScholarshipItems(req *entity.GetScholarshipItemsEntity) ([]*entity.ScholarshipItemEntity, error) {
	sil, err := s.scholarshipItemDao.GetList(model.DB.Self, req.ScholarshipId)
	if err != nil {
		return nil, err
	}

	scholarshipItems := make([]*entity.ScholarshipItemEntity, 0)
	for _, si := range sil {
		tmp := &entity.ScholarshipItemEntity{
			Id:            si.Id,
			Name:          si.Name,
			ScholarshipId: si.ScholarshipId,
		}

		scholarshipItems = append(scholarshipItems, tmp)
	}

	return scholarshipItems, nil
}

func (s *ScholarshipService) RemoveScholarshipItem(req *entity.RemoveScholarshipItemEntity) error {
	if err := s.scholarshipItemDao.DeleteByID(model.DB.Self, req.ScholarshipItemId); err != nil {
		return err
	}

	if err := s.applicationDao.DeleteByScholarshipItemId(model.DB.Self, int(req.ScholarshipItemId)); err != nil {
		return err
	}

	return nil
}
