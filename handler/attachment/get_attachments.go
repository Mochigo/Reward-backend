package attachment

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

func GetAttachments(c *gin.Context) {
	log.Info("GetAttachments called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	sid := c.Query("scholarship_id")
	if len(sid) == 0 {
		response.SendInternalServerError(c, errno.ErrRequiredParamsMissing, nil, "缺少scholarship_id")
		return
	}
	scholarshipIid, _ := strconv.Atoi(sid)

	entity := &entity.GetAttachmentsEntity{
		ScholarshipId: int64(scholarshipIid),
	}

	scholarshipService := service.NewScholarshipService(c)
	list, err := scholarshipService.GetAttachments(entity)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, list)
}
