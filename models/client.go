package models

type Client struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"not null"`

	Contacts []Contact `gorm:"foreignKey:ClientID"`
}

type CreateClientPayload struct {
	Name string `json:"name" binding:"required"`
}
