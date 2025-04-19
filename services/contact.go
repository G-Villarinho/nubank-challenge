package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/g-villarinho/nubank-challenge/repositories"
)

type ContactService interface {
	CreateContact(ctx context.Context, phone string, email string, clientId string) (*models.ContactResponse, error)
}

type contactService struct {
	di  *pkgs.Di
	clr repositories.ClientRepository
	ctr repositories.ContactRepository
}

func NewContactService(di *pkgs.Di) (ContactService, error) {
	clientRepository, err := pkgs.Invoke[repositories.ClientRepository](di)
	if err != nil {
		return nil, fmt.Errorf("invoke repositories.client: %w", err)
	}

	contactRepository, err := pkgs.Invoke[repositories.ContactRepository](di)
	if err != nil {
		return nil, fmt.Errorf("invoke repositories.contact: %w", err)
	}

	return &contactService{
		di:  di,
		clr: clientRepository,
		ctr: contactRepository,
	}, nil
}

func (c *contactService) CreateContact(ctx context.Context, phone string, email string, clientId string) (*models.ContactResponse, error) {
	client, err := c.clr.GetClientByID(ctx, clientId)
	if err != nil {
		return nil, fmt.Errorf("get client by id %s: %w", clientId, err)
	}

	if client == nil {
		return nil, models.ErrClientNotFound
	}

	contact := &models.Contact{
		Phone:    phone,
		Email:    email,
		ClientID: clientId,
	}

	if err := c.ctr.CreateContact(ctx, contact); err != nil {
		return nil, fmt.Errorf("create contact: %w", err)
	}

	return contact.ToContactResponse(), nil
}
