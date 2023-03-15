package scholarship

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
	"Reward/service/entity"
)

type GetScholarshipsResponse struct {
	Scholarships []*entity.ScholarshipEntity `json:"scholarships"`
	Total        int64                       `json:"total"`
}

func GetScholarships(c *gin.Context) {
	log.Info("GetScholarships called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	entity := &entity.GetScholarshipsEntity{
		Page:  page,
		Limit: limit,
	}

	scholarshipService := service.NewScholarshipService(c)
	list, total, err := scholarshipService.GetScholarships(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, GetScholarshipsResponse{
		Scholarships: list,
		Total:        total,
	})
}
