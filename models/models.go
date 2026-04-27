package models

import (
	"time"
	"gorm.io/gorm"
)

type NotificationTemplate struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"uniqueIndex;not null" json:"name"` // e.g. "welcome_email", "mfa_otp"
	Subject   string         `json:"subject"`
	Body      string         `gorm:"type:text;not null" json:"body"`
	Type      string         `json:"type"` // "email", "sms", "in_app"
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Notification struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     string         `gorm:"index" json:"user_id"`
	Type       string         `json:"type"` // "email", "sms", "in_app"
	Recipient  string         `json:"recipient"` // email address or phone number
	Subject    string         `json:"subject"`
	Content    string         `gorm:"type:text" json:"content"`
	Status     string         `json:"status"` // "pending", "sent", "failed"
	ReferenceID string        `json:"reference_id"` // external ID from ZeptoMail etc.
	Error      string         `json:"error,omitempty"`
	IsRead     bool           `gorm:"default:false" json:"is_read"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
