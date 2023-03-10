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

type LoginTeacherResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func LoginTeacher(c *gin.Context) {
	log.Info("LoginTeacher called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendBadRequest(c, errno.ErrBind, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	teacherService := service.NewTeacherService(c)
	ok, err := teacherService.VerifyTeacher(req.UID, req.Password)
	if err != nil {
		response.SendBadRequest(c, errno.ErrAuthFailed, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}
	if !ok {
		response.SendBadRequest(c, errno.ErrAuthFailed, nil, errors.New("没有对应用户").Error(), utils.GetUpFuncInfo(2))
		return
	}

	token, err := teacherService.Sign(req.UID)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrTokenGenerate, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	role, err := teacherService.GetTeacherRole(req.UID)
	if err != nil {
		response.SendInternalServerError(c, errno.ErrGetTeacherRole, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}

	response.SendResponse(c, errno.OK, &LoginTeacherResponse{
		Token: token,
		Role:  role,
	})
}
