package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"Reward/handler"
	"Reward/handler/attachment"
	"Reward/handler/certificate"
	"Reward/handler/login"
	"Reward/handler/sd"
	"Reward/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// global middleware
	g.Use(gin.Recovery())
	g.Use(middleware.RequestId())
	g.Use(middleware.Logging())
	g.Use(mw...)

	// 将本地的静态资源url，绑定到后者路径上
	g.Static(viper.GetString("file_storage"), "./statics")

	// 404 page.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	version := g.Group("api/v1.0")

	// login
	loginRouter := version.Group("/login")
	{
		loginRouter.POST("/student", login.LoginStudent)
		loginRouter.POST("/teacher", login.LoginTeacher)
	}

	// certificate
	certificateRouter := version.Group("/certificate")
	{
		certificateRouter.POST("/", certificate.AddCertificate)
		certificateRouter.GET("/list", certificate.GetCertificates)
		certificateRouter.PUT("/audit", certificate.AuditCertificate)
		certificateRouter.DELETE("/", certificate.RemoveCertificate)
	}

	// scholarship
	scholarshipRouter := version.Group("/scholarship")
	{
		scholarshipRouter.POST("/attchment", attachment.AddAttachment)
		scholarshipRouter.GET("/attchments", attachment.GetAttachments)
	}

	// upload
	uploadRouter := version.Group("/upload")
	{
		uploadRouter.POST("/", handler.UploadFile)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
