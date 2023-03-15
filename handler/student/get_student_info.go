package student

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common"
	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
)

func GetStudentInfo(c *gin.Context) {
	log.Info("GetCertificates called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	studentId := c.GetInt(common.TokenUserID)

	service := service.NewStudentService(c)
	stu, err := service.GetStudentInfo(int64(studentId))
	if err != nil {
		response.SendInternalServerError(c, errno.ErrBind, nil, err.Error())
	}

	response.SendResponse(c, nil, *stu)
}
