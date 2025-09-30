package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	SubID     string    `json:"sub_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
}
