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

type AddAttachmentRequest struct {
	ScholarshipId int64  `json:"scholarship_id"`
	Url           string `json:"url"`
}

func AddAttachment(c *gin.Context) {
	log.Info("AddAttachment called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AddAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.AddAttachmentEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	if err := scholarshipService.AddAttachment(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
