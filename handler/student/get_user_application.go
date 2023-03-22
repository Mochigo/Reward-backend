package student

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

type GetUserApplicationResponse struct {
	Applications []*entity.ApplicationEntity `json:"applications"`
	Total        int64                       `json:"total"`
}

func GetUserApplication(c *gin.Context) {
	log.Info("GetUserApplication called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	entity := &entity.GetUserApplicationEntity{
		Page:  page,
		Limit: limit,
	}

	studentService := service.NewStudentService(c)
	list, total, err := studentService.GetUserApplication(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, GetUserApplicationResponse{
		Applications: list,
		Total:        total,
	})
}
