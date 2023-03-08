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

type CreateScholarshipRequest struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type CreateScholarshipResponse struct {
	Id int64 `json:"scholarship_id"`
}

func CreateScholarship(c *gin.Context) {
	log.Info("CreateScholarship called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req CreateScholarshipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.CreateScholarshipEntity{}
	_ = utils.ConvertEntity(&req, entity)

	scholarshipService := service.NewScholarshipService(c)
	scholarshipId, err := scholarshipService.CreateScholarship(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}
	response.SendResponse(c, nil, CreateScholarshipResponse{
		Id: scholarshipId,
	})
}
