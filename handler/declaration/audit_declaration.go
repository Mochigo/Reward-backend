package declaration

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

type AuditDeclarationRequest struct {
	DeclarationId  int64  `json:"declaration_id"`
	RejectedReason string `json:"rejected_reason"` // 驳回理由
	Operation      string `json:"operation"`
}

func AuditDeclaration(c *gin.Context) {
	log.Info("AuditDeclaration called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AuditDeclarationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.AuditDeclarationEntity{}
	_ = utils.ConvertEntity(&req, entity)

	service := service.NewDeclarationService(c)
	if err := service.AuditDeclaration(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
