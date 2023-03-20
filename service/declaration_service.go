package service

import (
	"errors"

	"github.com/gin-gonic/gin"

	"Reward/common"
	"Reward/model"
	"Reward/service/entity"
)

type DeclarationService struct {
	ctx            *gin.Context
	declarationDao *model.DeclarationDao
}

func NewDeclarationService(ctx *gin.Context) *DeclarationService {
	return &DeclarationService{
		ctx:            ctx,
		declarationDao: model.GetDeclarationDao(),
	}
}

func (s *DeclarationService) AddDeclaration(req *entity.AddDeclarationEntity) error {
	return s.declarationDao.Create(model.DB.Self, &model.Declaration{
		ApplicationId: req.ApplicationId,
		StudentId:     int64(s.ctx.GetInt(common.TokenUserID)),
		Name:          req.Name,
		Level:         req.Level,
		Status:        common.StatusPROCESS,
		Url:           req.Url,
	})
}

// 后续考虑根据用户角色来返回地内容，如果是学生本人查看，则提供所有状态的荣誉查看，如果是直系老师查看只提供已经通过的荣誉的查看
func (s *DeclarationService) GetDeclarations(req *entity.GetDeclarationsEntity) ([]*entity.DeclarationEntity, error) {
	cl, err := s.declarationDao.GetDeclarationsByApplicationIdAndStatus(model.DB.Self, req.ApplicationId, common.StringEmpty)
	if err != nil {
		return nil, err
	}

	declarations := make([]*entity.DeclarationEntity, 0, len(cl))
	for _, c := range cl {
		tmp := &entity.DeclarationEntity{
			Id:             c.Id,
			ApplicationId:  c.ApplicationId,
			StudentId:      c.StudentId,
			Name:           c.Name,
			Level:          c.Level,
			Status:         c.Status,
			RejectedReason: c.RejectedReason,
			Url:            c.Url,
		}
		declarations = append(declarations, tmp)
	}

	return declarations, nil
}

func (s *DeclarationService) RemoveDeclaration(req *entity.RemoveDeclarationEntity) error {
	return s.declarationDao.Delete(model.DB.Self, req.DeclarationId)
}

func (s *DeclarationService) AuditDeclaration(req *entity.AuditDeclarationEntity) error {
	if req.Operation == common.OperationReject {
		return s.rejectDeclaration(req)
	}

	return s.approveDeclaration(req)
}

func (s *DeclarationService) rejectDeclaration(req *entity.AuditDeclarationEntity) error {
	if req.RejectedReason == common.StringEmpty {
		return errors.New("empty rejected reason")
	}

	updateMaps := make(map[string]interface{})
	updateMaps["status"] = common.StatusREJECTED
	updateMaps["rejected_reason"] = req.RejectedReason

	return s.declarationDao.Update(model.DB.Self, req.DeclarationId, updateMaps)
}

func (s *DeclarationService) approveDeclaration(req *entity.AuditDeclarationEntity) error {
	updateMaps := make(map[string]interface{})
	updateMaps["status"] = common.StatusAPPROVE

	return s.declarationDao.Update(model.DB.Self, req.DeclarationId, updateMaps)
}
