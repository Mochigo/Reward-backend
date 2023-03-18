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

func GetScholarshipInfo(c *gin.Context) {
	log.Info("GetScholarshipInfo called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	sid := c.Query("scholarship_id")
	if len(sid) == 0 {
		response.SendInternalServerError(c, errno.ErrRequiredParamsMissing, nil, "缺少scholarship_id")
		return
	}
	scholarshipId, _ := strconv.Atoi(sid)

	entity := &entity.GetScholarshipInfoEntity{
		ScholarshipId: int64(scholarshipId),
	}

	service := service.NewScholarshipService(c)
	scholarship, err := service.GetScholarshipInfo(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, *scholarship)
}
