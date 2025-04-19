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

type ClientRepository interface {
	CreateClient(ctx context.Context, client *models.Client) error
	GetClientsWithContact(ctx context.Context) ([]*models.Client, error)
	GetClientWitContactsByID(ctx context.Context, id string) (*models.Client, error)
	GetClientByID(ctx context.Context, id string) (*models.Client, error)
}

type clientRepository struct {
	di *pkgs.Di
	db *gorm.DB
}

func NewClientRepository(di *pkgs.Di) (ClientRepository, error) {
	db, err := pkgs.Invoke[*gorm.DB](di)
	if err != nil {
		return nil, fmt.Errorf("invoke gorm.DB: %w", err)
	}

	return &clientRepository{
		di: di,
		db: db,
	}, nil
}

func (c *clientRepository) CreateClient(ctx context.Context, client *models.Client) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate uuid: %w", err)
	}

	client.ID = id.String()
	client.CreatedAt = time.Now().UTC()

	if err := c.db.WithContext(ctx).Create(client).Error; err != nil {
		return err
	}

	return nil
}

func (c *clientRepository) GetClientsWithContact(ctx context.Context) ([]*models.Client, error) {
	var clients []*models.Client

	if err := c.db.WithContext(ctx).Preload("Contacts").Find(&clients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return clients, nil
}

func (c *clientRepository) GetClientWitContactsByID(ctx context.Context, id string) (*models.Client, error) {
	var client models.Client

	if err := c.db.WithContext(ctx).Preload("Contacts").First(&client, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &client, nil
}

func (c *clientRepository) GetClientByID(ctx context.Context, id string) (*models.Client, error) {
	var client models.Client

	if err := c.db.WithContext(ctx).First(&client, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &client, nil
}
