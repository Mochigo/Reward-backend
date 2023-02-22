package attachment

import (
	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetAttachmentsRequest struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

func GetAttachments(c *gin.Context) {
	log.Info("GetAttachments called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req GetAttachmentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.GetAttachmentsEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	list, err := scholarshipService.GetAttachments(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, list)
}
