package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kiyuu10/2fa-sys/controllers"
	"github.com/kiyuu10/2fa-sys/handlers"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/generate-send-otp", handlers.GenerateAndSendOTP)
		authGroup.POST("/verify-otp", handlers.VerifyOTP)
	}
}
