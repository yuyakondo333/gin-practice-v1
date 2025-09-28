package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	SubID     string    `gorm:"uniqueIndex;not null;column:sub_id"`
	Name      string    `gorm:"not null;"`
	AvatarURL string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"not null;default CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt time.Time `gorm:"not null;default CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt time.Time `gorm:"not null;default CURRENT_TIMESTAMP;column:deleted_at"`
}
