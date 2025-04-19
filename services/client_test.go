package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-villarinho/nubank-challenge/mocks"
	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateContact(t *testing.T) {
	ctx := context.Background()

	t.Run("should create contact successfully", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		client := &models.Client{ID: "client-123", Name: "Gabriel"}
		clientRepo.On("GetClientByID", ctx, "client-123").Return(client, nil)

		contactRepo.On("CreateContact", ctx, mock.MatchedBy(func(c *models.Contact) bool {
			return c.Phone == "123456789" && c.Email == "test@example.com" && c.ClientID == "client-123"
		})).Return(nil)

		resp, err := svc.CreateContact(ctx, "123456789", "test@example.com", "client-123")

		assert.NoError(t, err)
		assert.Equal(t, "123456789", resp.Phone)
		assert.Equal(t, "test@example.com", resp.Email)
	})

	t.Run("should return error if client not found", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.On("GetClientByID", ctx, "missing-client").Return(nil, nil)

		resp, err := svc.CreateContact(ctx, "123456789", "test@example.com", "missing-client")

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrClientNotFound)
	})

	t.Run("should return error if client repo fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.On("GetClientByID", ctx, "fail-client").Return(nil, errors.New("db error"))

		resp, err := svc.CreateContact(ctx, "123456789", "test@example.com", "fail-client")

		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "get client by id")
	})
}

func TestGetClientsWithContact(t *testing.T) {
	ctx := context.Background()

	t.Run("should return list of clients with contacts", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
		}

		mockClients := []*models.Client{
			{
				ID:        "client-1",
				Name:      "Gabriel",
				CreatedAt: time.Now(),
				Contacts: []models.Contact{
					{
						ID:        "contact-1",
						Phone:     "+5521999999999",
						Email:     "gabriel@gmail.com",
						ClientID:  "client-1",
						CreatedAt: time.Now(),
					},
				},
			},
		}

		clientRepo.On("GetClientsWithContact", ctx).Return(mockClients, nil)

		resp, err := svc.GetClientsWithContact(ctx)

		assert.NoError(t, err)
		assert.Len(t, resp, 1)
		assert.Equal(t, "Gabriel", resp[0].Name)
		assert.Equal(t, 1, len(resp[0].Contacts))
		assert.Equal(t, "gabriel@gmail.com", resp[0].Contacts[0].Email)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
		}

		clientRepo.On("GetClientsWithContact", ctx).Return(nil, errors.New("db failure"))

		resp, err := svc.GetClientsWithContact(ctx)

		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "get clients with contact")
	})
}

func TestGetClientContactsByID(t *testing.T) {
	ctx := context.Background()

	t.Run("should return contacts of a client", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		client := &models.Client{
			ID:   "client-123",
			Name: "Gabriel",
		}

		contacts := []*models.Contact{
			{
				ID:        "contact-1",
				Phone:     "+5521999999999",
				Email:     "gabriel@example.com",
				ClientID:  "client-123",
				CreatedAt: time.Now(),
			},
		}

		clientRepo.On("GetClientByID", ctx, "client-123").Return(client, nil)
		contactRepo.On("GetContactsByClientID", ctx, "client-123").Return(contacts, nil)

		resp, err := svc.GetClientContactsByID(ctx, "client-123")

		assert.NoError(t, err)
		assert.Len(t, resp, 1)
		assert.Equal(t, "gabriel@example.com", resp[0].Email)
	})

	t.Run("should return error if client not found", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.On("GetClientByID", ctx, "missing-client").Return(nil, nil)

		resp, err := svc.GetClientContactsByID(ctx, "missing-client")

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrClientNotFound)
	})

	t.Run("should return error if GetClientByID fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.On("GetClientByID", ctx, "client-123").Return(nil, errors.New("db error"))

		resp, err := svc.GetClientContactsByID(ctx, "client-123")

		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "get client by id")
	})

	t.Run("should return error if GetContactsByClientID fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		client := &models.Client{ID: "client-123", Name: "Gabriel"}

		clientRepo.On("GetClientByID", ctx, "client-123").Return(client, nil)
		contactRepo.On("GetContactsByClientID", ctx, "client-123").Return(nil, errors.New("repo error"))

		resp, err := svc.GetClientContactsByID(ctx, "client-123")

		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "get contacts by client id")
	})
}
