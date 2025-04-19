package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactRepository interface {
	CreateContact(ctx context.Context, contact *models.Contact) error
	GetContactsByClientID(ctx context.Context, clientID string) ([]*models.Contact, error)
}

type contactRepository struct {
	di *pkgs.Di
	db *gorm.DB
}

func NewContactRepository(di *pkgs.Di) (ContactRepository, error) {
	db, err := pkgs.Invoke[*gorm.DB](di)
	if err != nil {
		return nil, fmt.Errorf("invoke gorm.DB: %w", err)
	}

	return &contactRepository{
		di: di,
		db: db,
	}, nil
}

func (c *contactRepository) CreateContact(ctx context.Context, contact *models.Contact) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate uuid: %w", err)
	}

	contact.ID = id.String()
	contact.CreatedAt = time.Now().UTC()

	if err := c.db.WithContext(ctx).Create(contact).Error; err != nil {
		return err
	}

	return nil
}

func (c *contactRepository) GetContactsByClientID(ctx context.Context, clientID string) ([]*models.Contact, error) {
	var contacts []*models.Contact

	if err := c.db.WithContext(ctx).Where("client_id = ?", clientID).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}
