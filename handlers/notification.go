package handlers

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/eshdc/notification-service/config"
	"github.com/eshdc/notification-service/models"
	"github.com/eshdc/notification-service/utils"
	"github.com/gin-gonic/gin"
)

type SendNotificationRequest struct {
	UserID     string                 `json:"user_id"`
	Template   string                 `json:"template"` // e.g. "mfa_otp"
	Recipient  string                 `json:"recipient"` // email or phone
	Data       map[string]interface{} `json:"data"`      // dynamic data for template
	Name       string                 `json:"name"`      // recipient name
}

func ListTemplates(c *gin.Context) {
	var templates []models.NotificationTemplate
	if err := config.DB.Find(&templates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch templates"})
		return
	}
	c.JSON(http.StatusOK, templates)
}

func GetTemplate(c *gin.Context) {
	id := c.Param("id")
	var tpl models.NotificationTemplate
	if err := config.DB.First(&tpl, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}
	c.JSON(http.StatusOK, tpl)
}

func CreateTemplate(c *gin.Context) {
	var tpl models.NotificationTemplate
	if err := c.ShouldBindJSON(&tpl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&tpl).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create template"})
		return
	}
	c.JSON(http.StatusCreated, tpl)
}

func UpdateTemplate(c *gin.Context) {
	id := c.Param("id")
	var tpl models.NotificationTemplate
	if err := config.DB.First(&tpl, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}
	if err := c.ShouldBindJSON(&tpl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&tpl).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update template"})
		return
	}
	c.JSON(http.StatusOK, tpl)
}

func SendNotification(c *gin.Context) {
	var req SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Fetch template from DB
	var tpl models.NotificationTemplate
	if err := config.DB.Where("name = ?", req.Template).First(&tpl).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}

	// 2. Parse body with dynamic data
	tmpl, err := template.New(req.Template).Parse(tpl.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template body"})
		return
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, req.Data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute template"})
		return
	}

	// 3. Send via channel
	var refID string
	var sendErr error
	status := "sent"

	switch tpl.Type {
	case "email":
		refID, sendErr = utils.SendEmail(req.Recipient, req.Name, tpl.Subject, body.String())
	case "in_app":
		sendErr = utils.PublishNotification(req.UserID, gin.H{
			"subject": tpl.Subject,
			"body":    body.String(),
			"type":    "in_app",
		})
		refID = "redis_pub"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported notification type"})
		return
	}

	if sendErr != nil {
		status = "failed"
		refID = ""
	}

	// 4. Save audit log
	notification := models.Notification{
		UserID:      req.UserID,
		Type:        tpl.Type,
		Recipient:   req.Recipient,
		Subject:     tpl.Subject,
		Content:     body.String(),
		Status:      status,
		ReferenceID: refID,
	}
	if sendErr != nil {
		notification.Error = sendErr.Error()
	}
	config.DB.Create(&notification)

	if sendErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification", "details": sendErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "reference_id": refID})
}

func GetNotifications(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	var notifications []models.Notification
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Limit(50).Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func MarkAsRead(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Model(&models.Notification{}).Where("id = ?", id).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func MarkAllAsRead(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if err := config.DB.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark all as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
