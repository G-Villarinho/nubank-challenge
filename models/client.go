package models

import (
	"database/sql"
	"time"
)

type Client struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"not null"`

	Contacts  []Contact    `gorm:"foreignKey:ClientID"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`
}

type CreateClientPayload struct {
	Name string `json:"name" binding:"required"`
}
