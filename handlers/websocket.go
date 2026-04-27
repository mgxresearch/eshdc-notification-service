package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/eshdc/notification-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // In production, restrict this to allowed origins
	},
}

func HandleNotificationsWS(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to websocket: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("User %s connected to notifications", userID)

	// Subscribe to Redis channel for this user
	pubsub := utils.RedisClient.Subscribe(context.Background(), "notifications:"+userID)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			// Send message to websocket
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				log.Printf("Failed to send message to user %s: %v", userID, err)
				return
			}
		}
	}
}
