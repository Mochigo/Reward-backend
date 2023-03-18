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

type AddDeclarationRequest struct {
	Name          string `json:"name"`
	Level         string `json:"level"`
	Url           string `json:"url"`
	ApplicationId int64  `json:"application_id"`
}

func AddDeclaration(c *gin.Context) {
	log.Info("AddDeclaration called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AddDeclarationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.AddDeclarationEntity{}
	_ = utils.ConvertEntity(&req, entity)

	service := service.NewDeclarationService(c)
	if err := service.AddDeclaration(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
