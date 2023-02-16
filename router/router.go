package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Reward/handler/login"
	"Reward/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// global middleware
	g.Use(gin.Recovery())
	g.Use(middleware.RequestId())

	g.Use(mw...)

	// 404 page.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// login
	loginRouter := g.Group("/login")
	{
		loginRouter.POST("/student", login.LoginStudent)
		loginRouter.POST("/teacher", login.LoginTeacher)
	}

	return g
}
