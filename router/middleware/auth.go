package middleware

import (
	"github.com/gin-gonic/gin"

	"Reward/common/token"
)

// AuthMiddleware ... 认证中间件
func AuthMiddleware(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if len(tokenStr) == 0 {
		c.Abort()
		return
	}
	// Parse the json web token.
	payload, err := token.ParseToken(tokenStr)
	if err != nil {
		c.Abort()
		return
	}

	c.Set("userID", payload.UserID)

	c.Next()
}
