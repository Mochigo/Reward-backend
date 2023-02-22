package utils

import (
	"github.com/gin-gonic/gin"

	"Reward/common"
)

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return common.StringEmpty
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return common.StringEmpty
}
