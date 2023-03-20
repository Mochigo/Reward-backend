package teacher

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

type GetStudentsDeclarationsResponse struct {
	Declarations []*entity.DeclarationEntity `json:"declarations"`
	Total        int64                       `json:"total"`
}

func GetStudentsDeclarations(c *gin.Context) {
	log.Info("GetStudentsDeclarations called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	entity := &entity.GetStudentsDeclarationsEntity{
		Page:  page,
		Limit: limit,
	}
	service := service.NewTeacherService(c)
	list, total, err := service.GetStudentsDeclarations(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, GetStudentsDeclarationsResponse{
		Declarations: list,
		Total:        total,
	})
}
