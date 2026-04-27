package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eshdc/notification-service/config"
	"github.com/eshdc/notification-service/handlers"
	"github.com/eshdc/notification-service/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB & Redis
	config.InitDB()
	utils.InitRedis()

	if os.Getenv("SEED") == "true" {
		config.SeedTemplates()
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := gin.Default()

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
			"service": "notification-service",
		})
	})

	// API Routes
	api := r.Group("/api/v1")
	{
		notifications := api.Group("/notifications")
		{
			notifications.GET("", handlers.GetNotifications)
			notifications.POST("/read/:id", handlers.MarkAsRead)
			notifications.POST("/read-all", handlers.MarkAllAsRead)
			notifications.POST("/send", handlers.SendNotification)
			notifications.GET("/ws", handlers.HandleNotificationsWS)
		}

		templates := api.Group("/templates")
		{
			templates.GET("", handlers.ListTemplates)
			templates.GET("/:id", handlers.GetTemplate)
			templates.POST("", handlers.CreateTemplate)
			templates.PUT("/:id", handlers.UpdateTemplate)
		}
	}

	log.Printf("Notification Service starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
