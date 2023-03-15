package middleware

import (
	"github.com/gin-gonic/gin"

	"Reward/common"
	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/token"
)

// AuthMiddleware ... 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			c.Abort()
			return
		}
		// Parse the json web token.
		payload, err := token.ParseToken(tokenStr)
		if err != nil {
			response.SendUnauthorized(c, errno.ErrTokenInvalid, nil, err.Error())
			c.Abort()
			return
		}

		c.Set(common.TokenUserID, payload.UserID)
		c.Set(common.TokenCollegeID, payload.CollegeId)

		c.Next()
	}
}
