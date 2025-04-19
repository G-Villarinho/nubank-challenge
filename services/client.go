package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/g-villarinho/nubank-challenge/repositories"
)

type ClientService interface {
	CreateClient(ctx context.Context, name string, contacts []*models.Contact) (*models.ClientResponse, error)
	GetClientsWithContact(ctx context.Context) ([]*models.ClientResponse, error)
	GetClientContactsByID(ctx context.Context, id string) ([]*models.ContactResponse, error)
}

type clientService struct {
	di  *pkgs.Di
	clr repositories.ClientRepository
	ctr repositories.ContactRepository
}

func NewClientService(di *pkgs.Di) (ClientService, error) {
	clientRepository, err := pkgs.Invoke[repositories.ClientRepository](di)
	if err != nil {
		return nil, fmt.Errorf("invoke repositories.Client: %w", err)
	}

	contactRepository, err := pkgs.Invoke[repositories.ContactRepository](di)
	if err != nil {
		return nil, fmt.Errorf("invoke repositories.Contact: %w", err)
	}

	return &clientService{
		di:  di,
		clr: clientRepository,
		ctr: contactRepository,
	}, nil
}

func (c *clientService) CreateClient(ctx context.Context, name string, contacts []*models.Contact) (*models.ClientResponse, error) {
	client := &models.Client{
		Name: name,
	}

	if err := c.clr.CreateClient(ctx, client); err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	if len(contacts) > 0 {
		for i := range contacts {
			contacts[i].ClientID = client.ID
		}

		if err := c.ctr.CreateContacts(ctx, contacts); err != nil {
			return nil, fmt.Errorf("create contacts: %w", err)
		}
	}

	clientResponse := &models.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		CreatedAt: client.CreatedAt,
	}

	clientResponse.Contacts = make([]models.ContactResponse, len(contacts))
	for i, contact := range contacts {
		clientResponse.Contacts[i] = models.ContactResponse{
			ID:        contact.ID,
			Phone:     contact.Phone,
			Email:     contact.Email,
			CreatedAt: contact.CreatedAt,
		}
	}

	return clientResponse, nil
}

func (c *clientService) GetClientsWithContact(ctx context.Context) ([]*models.ClientResponse, error) {
	clients, err := c.clr.GetClientsWithContact(ctx)
	if err != nil {
		return nil, fmt.Errorf("get clients with contact: %w", err)
	}

	clientResponses := make([]*models.ClientResponse, 0, len(clients))

	for _, client := range clients {
		clientResponses = append(clientResponses, &models.ClientResponse{
			ID:        client.ID,
			Name:      client.Name,
			CreatedAt: client.CreatedAt,
			Contacts:  toContactResponses(client.Contacts),
		})
	}

	return clientResponses, nil
}

func (c *clientService) GetClientContactsByID(ctx context.Context, id string) ([]*models.ContactResponse, error) {
	client, err := c.clr.GetClientByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get client by id %s: %w", id, err)
	}

	if client == nil {
		return nil, models.ErrClientNotFound
	}

	contacts, err := c.ctr.GetContactsByClientID(ctx, client.ID)
	if err != nil {
		return nil, fmt.Errorf("get contacts by client id %s: %w", client.ID, err)
	}

	contactsResponse := make([]*models.ContactResponse, 0, len(contacts))
	for _, contact := range contacts {
		contactsResponse = append(contactsResponse, &models.ContactResponse{
			ID:        contact.ID,
			Phone:     contact.Phone,
			Email:     contact.Email,
			CreatedAt: contact.CreatedAt,
		})
	}

	return contactsResponse, nil
}

func toContactResponses(contacts []models.Contact) []models.ContactResponse {
	contactResponses := make([]models.ContactResponse, len(contacts))
	for i, contact := range contacts {
		contactResponses[i] = models.ContactResponse{
			ID:    contact.ID,
			Email: contact.Email,
			Phone: contact.Phone,
		}
	}
	return contactResponses
}
