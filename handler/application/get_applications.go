package application

import (
	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
	"Reward/service/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetItemApplicationsResponse struct {
	Applications []*entity.ItemApplicationEntity `json:"applications"`
	Total        int64                           `json:"total"`
}

func GetItemApplications(c *gin.Context) {
	log.Info("GetItemApplications called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	scholarshipItemId := c.Query("scholarship_item_id")
	if len(scholarshipItemId) == 0 {
		response.SendInternalServerError(c, errno.ErrRequiredParamsMissing, nil, "缺少scholarship_item_id")
		return
	}

	id, _ := strconv.Atoi(scholarshipItemId)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	entity := &entity.GetItemApplicationsEntity{
		Page:              page,
		Limit:             limit,
		ScholarshipItemId: int64(id),
	}

	applicationService := service.NewApplicationService(c)
	list, total, err := applicationService.GetItemApplications(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, GetItemApplicationsResponse{
		Applications: list,
		Total:        total,
	})
}
