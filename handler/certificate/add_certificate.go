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

type AddCertificateRequest struct {
	Name          string `json:"name"`
	Level         string `json:"level"`
	Url           string `json:"url"`
	ApplicationId int64  `json:"application_id"`
}

type AddCertificateResponse struct {
	CertificateID int64 `json:"certificate_id"`
}

func AddCertificate(c *gin.Context) {
	log.Info("AddCertificate called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AddCertificateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.AddCertificateEntity{}
	_ = utils.ConvertEntity(&req, entity)

	cs := service.NewCertificateService(c)
	cid, err := cs.AddCertificate(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, AddCertificateResponse{cid})
}
