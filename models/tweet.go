package models

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt time.Time `gorm:"not null;default CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"` //nilを許可
}
