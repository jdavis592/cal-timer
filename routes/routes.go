package routes

import (
	"cal-timer/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/auth/google", controllers.AuthGoogle)
		api.GET("/auth/callback", controllers.AuthCallback)
	}
}
