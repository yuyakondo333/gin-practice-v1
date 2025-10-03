package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	SubID     string
	Name      string
	AvatarURL string
}

var ErrUserNotFound = errors.New("user not found")

// 将来ここに業務ルール（Validate(), ChangeName(), etc.）を生やせる
