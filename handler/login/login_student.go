package login

import (
	"errors"

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
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	studentService := service.NewStudentService(c)
	ok, err := studentService.VerifyStudent(req.UID, req.Password)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrAuthFailed, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}
	if !ok {
		response.SendInternalServerError(c, errno.ErrAuthFailed, nil, errors.New("没有对应用户").Error(), utils.GetUpFuncInfo(2))
		return
	}

	token, err := studentService.Sign(req.UID)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrTokenGenerate, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, errno.OK, &LoginResponse{
		Token: token,
	})
}
