package certificate

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

type GetCertificatesRequest struct {
	ApplicationId int64 `json:"application_id"`
}

func GetCertificates(c *gin.Context) {
	log.Info("GetCertificates called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req GetCertificatesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	entity := &entity.GetCertificatesEntity{}
	_ = utils.ConvertEntity(&req, entity)

	cs := service.NewCertificateService(c)
	list, err := cs.GetCertificates(entity)
	if err != nil {
		//TODO 修改所有的Errbind
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, nil, list)
}
