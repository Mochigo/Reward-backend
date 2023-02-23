package certificate

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
	"Reward/service/entity"
)

type AuditCertificateRequest struct {
	CertificateId  int64  `json:"certificate_id"`
	RejectedReason string `json:"rejected_reason"` // 驳回理由
	Operation      string `json:"operation"`
}

func AuditCertificate(c *gin.Context) {
	log.Info("AuditCertificate called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AuditCertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.AuditCertificateEntity{}
	_ = utils.ConvertEntity(&req, entity)

	cs := service.NewCertificateService(c)
	if err := cs.AuditCertificate(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, nil)
}
