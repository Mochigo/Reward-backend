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

type RemoveScholarshipItemRequest struct {
	ScholarshipItemId int64 `json:"scholarship_item_id"`
}

func RemoveScholarshipItem(c *gin.Context) {
	log.Info("RemoveScholarshipItem called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req RemoveScholarshipItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.RemoveScholarshipItemEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	if err := scholarshipService.RemoveScholarshipItem(entity); err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, nil)
}
