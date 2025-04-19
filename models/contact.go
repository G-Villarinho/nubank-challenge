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

type ContactResponse struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *Contact) ToContactResponse() *ContactResponse {
	return &ContactResponse{
		ID:        c.ID,
		Phone:     c.Phone,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}
}

func ToContacts(payloads []CreateContactPayload) []*Contact {
	contacts := make([]*Contact, len(payloads))
	for i, p := range payloads {
		contacts[i] = &Contact{
			Phone: p.Phone,
			Email: p.Email,
		}
	}
	return contacts
}

func ToContactResponses(contacts []Contact) []*ContactResponse {
	res := make([]*ContactResponse, len(contacts))
	for i, c := range contacts {
		res[i] = &ContactResponse{
			ID:        c.ID,
			Phone:     c.Phone,
			Email:     c.Email,
			CreatedAt: c.CreatedAt,
		}
	}
	return res
}
