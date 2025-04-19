package models

import (
	"database/sql"
	"time"
)

type Contact struct {
	ID    string `gorm:"type:uuid;primaryKey"`
	Phone string `gorm:"not null"`
	Email string `gorm:"not null"`

	ClientID string `gorm:"type:uuid;not null"`
	Client   Client `gorm:"foreignKey:ClientID"`

	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`
}

type CreateContactPayload struct {
	Phone    string `json:"phone" binding:"required,e164"`
	Email    string `json:"email" binding:"required,email"`
	ClientID string `json:"clientId" binding:"required,uuid"`
}
