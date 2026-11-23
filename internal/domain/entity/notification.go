package entity

import (
	"time"

	"github.com/google/uuid"
)

// Notification merepresentasikan notifikasi in-platform (PRD M8).
type Notification struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID            uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index:idx_notifications_user_id"`
	User              User       `json:"user" gorm:"foreignKey:UserID"`
	Type              string     `json:"type" gorm:"type:varchar(50);not null"`
	Title             string     `json:"title" gorm:"type:varchar(255);not null"`
	Message           string     `json:"message" gorm:"type:text;not null"`
	RelatedEntityID   *uuid.UUID `json:"related_entity_id,omitempty" gorm:"type:uuid"`
	RelatedEntityType string     `json:"related_entity_type" gorm:"type:varchar(50)"`
	IsRead            bool       `json:"is_read" gorm:"type:boolean;default:false;index:idx_notifications_is_read"`
	CreatedAt         time.Time  `json:"created_at" gorm:"type:timestamptz;default:now();index:idx_notifications_created_at"`
}

// TableName mengembalikan nama tabel di PostgreSQL.
func (Notification) TableName() string {
	return "notifications"
}
