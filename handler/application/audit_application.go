package application

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

type AuditApplicationRequest struct {
	ApplicationId int64 `json:"application_id"`
}

func AuditApplication(c *gin.Context) {
	log.Info("AuditDeclaration called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AuditApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.AuditApplicationEntity{}
	_ = utils.ConvertEntity(&req, entity)

	service := service.NewApplicationService(c)
	if err := service.AuditApplication(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
