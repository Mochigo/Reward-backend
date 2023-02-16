package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/utils"
	"Reward/log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func sendError(c *gin.Context, status int, err error, data interface{}, cause string, pos string) {
	code, message := errno.DecodeErr(err)
	log.Info(message,
		zap.String("X-Request-Id", utils.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("pos", pos),
	)

	c.JSON(status, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string, pos string) {
	sendError(c, http.StatusBadRequest, err, data, cause, pos)
}

func SendUnauthorized(c *gin.Context, err error, data interface{}, cause string, pos string) {
	sendError(c, http.StatusUnauthorized, err, data, cause, pos)
}

func SendForbidden(c *gin.Context, err error, data interface{}, cause string, pos string) {
	sendError(c, http.StatusForbidden, err, data, cause, pos)
}

func SendNotFound(c *gin.Context, err error, data interface{}, cause string, pos string) {
	sendError(c, http.StatusNotFound, err, data, cause, pos)
}

func SendInternalServerError(c *gin.Context, err error, data interface{}, cause string, pos string) {
	sendError(c, http.StatusInternalServerError, err, data, cause, pos)
}
