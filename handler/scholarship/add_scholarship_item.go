package scholarship

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

type AddScholarshipItemRequest struct {
	ScholarshipId int64  `json:"scholarship_id"`
	Name          string `json:"name"`
}

func AddScholarshipItem(c *gin.Context) {
	log.Info("AddScholarshipItem called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req AddScholarshipItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	entity := &entity.AddScholarshipItemEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	if err := scholarshipService.AddScholarshipItem(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, nil)
}
