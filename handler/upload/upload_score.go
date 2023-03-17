package upload

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
	"Reward/service/parse"
)

func UploadScore(c *gin.Context) {
	log.Info("UploadScore function called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	file, err := c.FormFile("file")
	if err != nil {
		response.SendInternalServerError(c, errno.ErrFileNotFound, nil, err.Error())
		return
	}

	ps := parse.NewParseService(c)
	students, err := ps.Parse(file)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrParsing, nil, err.Error())
		return
	}

	service := service.NewStudentService(c)
	resp, err := service.UploadScore(students)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrUploadScore, nil, err.Error())
		return
	}

	response.SendResponse(c, errno.OK, *resp)
}
