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

func GetScholarshipItems(c *gin.Context) {
	log.Info("GetScholarshipItems called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	sid := c.Query("scholarship_id")
	if len(sid) == 0 {
		response.SendInternalServerError(c, errno.ErrRequiredParamsMissing, nil, "缺少scholarship_id")
		return
	}
	scholarshipIid, _ := strconv.Atoi(sid)

	entity := &entity.GetScholarshipItemsEntity{
		ScholarshipId: int64(scholarshipIid),
	}

	scholarshipService := service.NewScholarshipService(c)
	list, err := scholarshipService.GetScholarshipItems(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, list)
}
