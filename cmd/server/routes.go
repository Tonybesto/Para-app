package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"paramount_school/internal/api"
	"paramount_school/internal/middleware"
	"paramount_school/internal/ports"
	"time"
)

//SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{
		r.GET("/ping", handler.PingHandler)
		r.POST("/register", handler.SignUpHandler)
		r.POST("/student_login", handler.LoginStudentHandler)
		r.POST("/login/phone", handler.PhoneLoginStudentHandler)
		r.GET("/get_high_schools", handler.GetHighSchools)
		r.POST("/create_high_schools", handler.CreateHighSchool)
	}

	// authorizeStudent authorizes all authorized student handlers
	authorizeStudent := r.Group("/student")
	authorizeStudent.Use(middleware.AuthorizeStudent(repository.FindStudentByEmail, repository.TokenInBlacklist))
	{
		authorizeStudent.PUT("update_profile", handler.StudentUpdateHandler)
		authorizeStudent.GET("get_profile", handler.StudentProfileHandler)

	}

	authorizeMobileLogin := r.Group("/mobile")
	authorizeMobileLogin.Use(middleware.AuthorizeMobileLogin(repository.FindStudentByPhone, repository.TokenInBlacklist))
	{
		authorizeMobileLogin.POST("/phone/verify", handler.VerifyPhoneLoginStudentHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
