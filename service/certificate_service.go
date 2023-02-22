package service

import (
	"Reward/common"
	"Reward/model"
	"Reward/service/entity"
	"errors"

	"github.com/gin-gonic/gin"
)

type CertificateService struct {
	Ctx            *gin.Context
	certificateDao *model.CertificateDao
}

func NewCertificateService(ctx *gin.Context) *CertificateService {
	return &CertificateService{
		Ctx:            ctx,
		certificateDao: model.GetCertificateDao(),
	}
}

func (s *CertificateService) AddCertificate(req *entity.AddCertificateEntity) (int64, error) {
	c := &model.Certificate{
		ApplicationId: req.ApplicationId,
		Name:          req.Name,
		Level:         req.Level,
		Status:        common.StatusPROCESS,
		Url:           req.Url,
	}
	err := s.certificateDao.Create(model.DB.Self, c)
	if err != nil {
		return 0, err
	}

	return c.Id, nil
}

// 后续考虑根据用户角色来返回地内容，如果是学生本人查看，则提供所有状态的荣誉查看，如果是直系老师查看只提供已经通过的荣誉的查看
func (s *CertificateService) GetCertificates(req *entity.GetCertificatesEntity) ([]*entity.CertificateEntity, error) {
	cl, err := s.certificateDao.GetCertificatesByApplicationIdAndStatus(model.DB.Self, req.ApplicationId, common.StringEmpty)
	if err != nil {
		return nil, err
	}

	certificates := make([]*entity.CertificateEntity, 0, len(cl))
	for _, c := range cl {
		tmp := &entity.CertificateEntity{
			ApplicationId:  c.ApplicationId,
			Name:           c.Name,
			Level:          c.Level,
			Status:         c.Status,
			RejectedReason: c.RejectedReason,
			Url:            c.Url,
		}
		certificates = append(certificates, tmp)
	}

	return certificates, nil
}

func (s *CertificateService) RemoveCertificate(req *entity.RemoveCertificateEntity) error {
	return s.certificateDao.Delete(model.DB.Self, req.CertificateId)
}

func (s *CertificateService) AuditCertificate(req *entity.AuditCertificateEntity) error {
	if req.Operation == common.OperationReject {
		return s.rejectCertificate(req)
	}

	return s.approveCertificate(req)
}

func (s *CertificateService) rejectCertificate(req *entity.AuditCertificateEntity) error {
	if req.RejectedReason == common.StringEmpty {
		return errors.New("empty rejected reason")
	}

	updateMaps := make(map[string]interface{})
	updateMaps["status"] = common.StatusREJECTED
	updateMaps["rejected_reason"] = req.RejectedReason

	return s.certificateDao.Update(model.DB.Self, req.CertificateId, updateMaps)
}

func (s *CertificateService) approveCertificate(req *entity.AuditCertificateEntity) error {
	updateMaps := make(map[string]interface{})
	updateMaps["status"] = common.StatusAPPROVE

	return s.certificateDao.Update(model.DB.Self, req.CertificateId, updateMaps)
}
