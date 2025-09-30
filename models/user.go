package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubID     string         `gorm:"uniqueIndex;not null;column:sub_id"`
	Name      string         `gorm:"not null;"`
	AvatarURL string         `gorm:"not null;default:''"`
	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
