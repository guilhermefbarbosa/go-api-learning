package utils

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"default:uuid_generate_v4() primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type WeakModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}
