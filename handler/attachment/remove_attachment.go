package attachment

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

type RemoveAttachmentRequest struct {
	AttachmentId int64 `json:"attachment_id"`
}

func RemoveAttachment(c *gin.Context) {
	log.Info("RemoveAttachment called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req RemoveAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.AddAttachmentEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	if err := scholarshipService.AddAttachment(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrDeleteAttachment, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, nil)
}
