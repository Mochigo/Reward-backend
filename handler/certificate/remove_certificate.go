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

type RemoveCertificateRequest struct {
	CertificateId int64 `json:"certificate_id"`
}

func RemoveCertificate(c *gin.Context) {
	log.Info("RemoveCertificate called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req RemoveCertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.RemoveCertificateEntity{}
	_ = utils.ConvertEntity(&req, entity)

	cs := service.NewCertificateService(c)
	if err := cs.RemoveCertificate(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, nil)
}
