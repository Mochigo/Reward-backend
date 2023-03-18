package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"Reward/handler/application"
	"Reward/handler/attachment"
	"Reward/handler/declaration"
	"Reward/handler/login"
	"Reward/handler/scholarship"
	"Reward/handler/sd"
	"Reward/handler/student"
	"Reward/handler/upload"
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

	// declaration
	declarationRouter := version.Group("/declaration")
	{
		declarationRouter.POST("", declaration.AddDeclaration)
		declarationRouter.GET("/list", declaration.GetDeclarations)
		declarationRouter.PUT("/audit", declaration.AuditDeclaration)
		declarationRouter.DELETE("", declaration.RemoveDeclaration)
	}

	// scholarship
	scholarshipRouter := version.Group("/scholarship", middleware.AuthMiddleware())
	{
		scholarshipRouter.POST("/attachment", attachment.AddAttachment)
		scholarshipRouter.DELETE("/attachment", attachment.RemoveAttachment)
		scholarshipRouter.GET("/attachments", attachment.GetAttachments)

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

	studentRouter := version.Group("/student", middleware.AuthMiddleware())
	{
		studentRouter.GET("", student.GetStudentInfo)
	}

	// upload
	uploadRouter := version.Group("/upload", middleware.AuthMiddleware())
	{
		uploadRouter.POST("", upload.UploadFile)
		uploadRouter.POST("/score", upload.UploadScore)
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
