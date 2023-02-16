package login

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
	"Reward/service"
)

func LoginTeacher(c *gin.Context) {
	log.Info("LoginTeacher called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	teacherService := service.NewTeacherService(c)
	if ok, err := teacherService.VerifyTeacher(req.UID, req.Password); !ok {
		response.SendBadRequest(c, errno.ErrAuthFailed, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	token, err := teacherService.Sign(req.UID)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrTokenGenerate, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, errno.OK, &LoginResponse{
		Token: token,
	})
}
