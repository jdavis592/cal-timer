package main

import (
	"cal-timer/auth"
	"cal-timer/config"
	"cal-timer/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	auth.InitOAuthConfig()

	router := gin.Default()
	routes.SetupRoutes(router)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
