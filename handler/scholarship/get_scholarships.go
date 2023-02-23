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

type GetScholarshipsRequest struct {
	CollegeId int64 `json:"college_id"`
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
}

func GetScholarships(c *gin.Context) {
	log.Info("GetScholarships called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req GetScholarshipsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.GetScholarshipsEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	list, err := scholarshipService.GetScholarships(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, list)
}
