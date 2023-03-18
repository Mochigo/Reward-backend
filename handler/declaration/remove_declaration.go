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

type RemoveDeclarationRequest struct {
	DeclarationId int64 `json:"declaration_id"`
}

func RemoveDeclaration(c *gin.Context) {
	log.Info("RemoveCertificate called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req RemoveDeclarationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.RemoveDeclarationEntity{}
	_ = utils.ConvertEntity(&req, entity)

	service := service.NewDeclarationService(c)
	if err := service.RemoveDeclaration(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
