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

type LoginRequest struct {
	UID      string `json:"uid"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginStudent(c *gin.Context) {
	log.Info("LoginStudent called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	studentService := service.NewStudentService(c)
	ok, err := studentService.VerifyStudent(req.UID, req.Password)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrAuthFailed, nil, err.Error())
		return
	}
	if !ok {
		response.SendInternalServerError(c, errno.ErrAuthFailed, nil, "账户不存在或者用户名密码错误")
		return
	}

	token, err := studentService.Sign(req.UID)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrTokenGenerate, nil, err.Error())
		return
	}

	response.SendResponse(c, errno.OK, &LoginResponse{
		Token: token,
	})
}
