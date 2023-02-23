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

type GetScholarshipItemsRequest struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

func GetScholarshipItems(c *gin.Context) {
	log.Info("GetScholarshipItems called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req GetScholarshipItemsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}
	entity := &entity.GetScholarshipItemsEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	list, err := scholarshipService.GetScholarshipItems(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, list)
}
