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

func TestCreateClient(t *testing.T) {
	ctx := context.Background()

	t.Run("should create client with contacts successfully", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		contacts := []*models.Contact{
			{
				ID:    "contact-1",
				Phone: "+5521999999999",
				Email: "gabriel@gmail.com",
			},
		}

		clientRepo.
			On("CreateClient", ctx, mock.MatchedBy(func(c *models.Client) bool {
				c.ID = "client-123"
				c.CreatedAt = time.Now()
				return c.Name == "Gabriel"
			})).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*models.Client)
				arg.ID = "client-123"
				arg.CreatedAt = time.Now()
			}).
			Return(nil)

		contactRepo.
			On("CreateContacts", ctx, mock.Anything).
			Return(nil)

		resp, err := svc.CreateClient(ctx, "Gabriel", contacts)

		assert.NoError(t, err)
		assert.Equal(t, "Gabriel", resp.Name)
		assert.Len(t, resp.Contacts, 1)
		assert.Equal(t, "+5521999999999", resp.Contacts[0].Phone)
	})

	t.Run("should create client without contacts", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.
			On("CreateClient", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*models.Client)
				arg.ID = "client-123"
				arg.CreatedAt = time.Now()
			}).
			Return(nil)

		resp, err := svc.CreateClient(ctx, "Sem Contato", nil)

		assert.NoError(t, err)
		assert.Equal(t, "Sem Contato", resp.Name)
		assert.Empty(t, resp.Contacts)
	})

	t.Run("should return error if client creation fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.
			On("CreateClient", ctx, mock.Anything).
			Return(errors.New("erro no banco"))

		resp, err := svc.CreateClient(ctx, "Gabriel", nil)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "create client")
	})

	t.Run("should return error if creating contacts fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		svc := &clientService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		contacts := []*models.Contact{
			{
				ID:    "c1",
				Phone: "+5521999999999",
				Email: "g@a.com",
			},
		}

		clientRepo.
			On("CreateClient", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(*models.Client)
				arg.ID = "client-123"
				arg.CreatedAt = time.Now()
			}).
			Return(nil)

		contactRepo.
			On("CreateContacts", ctx, mock.Anything).
			Return(errors.New("erro contatos"))

		resp, err := svc.CreateClient(ctx, "Gabriel", contacts)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "create contacts")
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
