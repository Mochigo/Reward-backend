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

type CreateApplicationRequest struct {
	ScholarshipItemId int64  `json:"scholarship_item_id"` // 奖学金子项id
	ScholarshipId     int64  `json:"scholarship_id"`      // 奖学金id
	Deadline          string `json:"deadline"`
}

func CreateApplication(c *gin.Context) {
	log.Info("CreateApplication called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.CreateApplicationEntity{}
	_ = utils.ConvertEntity(&req, entity)

	applicationService := service.NewApplicationService(c)
	if err := applicationService.CreateApplication(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
