package certificate

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

func GetCertificates(c *gin.Context) {
	log.Info("GetCertificates called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	aid := c.Query("application_id")
	if len(aid) == 0 {
		response.SendInternalServerError(c, errno.ErrRequiredParamsMissing, nil, "缺少application_id")
		return
	}
	applicationId, _ := strconv.Atoi(aid)

	entity := &entity.GetCertificatesEntity{
		ApplicationId: int64(applicationId),
	}

	cs := service.NewCertificateService(c)
	list, err := cs.GetCertificates(entity)
	if err != nil {
		//TODO 修改所有的Errbind
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
		return
	}

	response.SendResponse(c, nil, list)
}
