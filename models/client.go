package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

type Client struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"not null"`

	Contacts  []Contact    `gorm:"foreignKey:ClientID"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`
}

type CreateClientPayload struct {
	Name     string                 `json:"name" binding:"required" example:"Gabriel Villarinho"`
	Contacts []CreateContactPayload `json:"contacts" binding:"required"`
}

type ClientResponse struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"created_at"`
	Contacts  []*ContactResponse `json:"contacts"`
}

func (c *Client) ToClientResponse() *ClientResponse {
	return &ClientResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		Contacts:  ToContactResponses(c.Contacts),
	}
}
