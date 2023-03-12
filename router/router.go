package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"Reward/handler"
	"Reward/handler/application"
	"Reward/handler/attachment"
	"Reward/handler/certificate"
	"Reward/handler/login"
	"Reward/handler/scholarship"
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

	version := g.Group("/api/v1.0")

	// login
	loginRouter := version.Group("/login")
	{
		loginRouter.POST("/student", login.LoginStudent)
		loginRouter.POST("/teacher", login.LoginTeacher)
	}

	// certificate
	certificateRouter := version.Group("/certificate")
	{
		certificateRouter.POST("", certificate.AddCertificate)
		certificateRouter.GET("/list", certificate.GetCertificates)
		certificateRouter.PUT("/audit", certificate.AuditCertificate)
		certificateRouter.DELETE("", certificate.RemoveCertificate)
	}

	// scholarship
	scholarshipRouter := version.Group("/scholarship", middleware.AuthMiddleware())
	{
		scholarshipRouter.POST("/attchment", attachment.AddAttachment)
		scholarshipRouter.DELETE("/attchment", attachment.RemoveAttachment)
		scholarshipRouter.GET("/attchments", attachment.GetAttachments)

		scholarshipRouter.POST("", scholarship.CreateScholarship)
		scholarshipRouter.GET("/list", scholarship.GetScholarships)

		scholarshipRouter.POST("/item", scholarship.AddScholarshipItem)
		scholarshipRouter.GET("/items/list", scholarship.GetScholarshipItems)
		scholarshipRouter.DELETE("/item", scholarship.RemoveScholarshipItem)
	}

	applicationRouter := version.Group("/application", middleware.AuthMiddleware())
	{
		applicationRouter.POST("", application.CreateApplication)
		applicationRouter.GET("/list", application.GetUserApplication)
	}

	// upload
	uploadRouter := version.Group("/upload")
	{
		uploadRouter.POST("", handler.UploadFile)
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
