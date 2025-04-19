package services

import (
	"context"
	"errors"
	"testing"

	"github.com/g-villarinho/nubank-challenge/mocks"
	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContactService_CreateContact(t *testing.T) {
	ctx := context.Background()

	t.Run("should create contact successfully", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		service := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		client := &models.Client{ID: "client-123", Name: "Gabriel"}

		clientRepo.
			On("GetClientByID", ctx, "client-123").
			Return(client, nil)

		contactRepo.
			On("CreateContact", ctx, mock.MatchedBy(func(c *models.Contact) bool {
				return c.Phone == "123456789" && c.Email == "test@example.com" && c.ClientID == "client-123"
			})).
			Return(nil)

		result, err := service.CreateContact(ctx, "123456789", "test@example.com", "client-123")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "123456789", result.Phone)
		assert.Equal(t, "test@example.com", result.Email)
	})

	t.Run("should return error if client not found", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		service := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.
			On("GetClientByID", ctx, "missing-client").
			Return(nil, nil)

		result, err := service.CreateContact(ctx, "123456789", "test@example.com", "missing-client")

		assert.ErrorIs(t, err, models.ErrClientNotFound)
		assert.Nil(t, result)
	})

	t.Run("should return error when client repo fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		service := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		clientRepo.
			On("GetClientByID", ctx, "client-error").
			Return(nil, errors.New("db failure"))

		result, err := service.CreateContact(ctx, "123456789", "test@example.com", "client-error")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "get client by id")
	})

	t.Run("should return error when creating contact fails", func(t *testing.T) {
		clientRepo := new(mocks.ClientRepositoryMock)
		contactRepo := new(mocks.ContactRepositoryMock)

		service := &contactService{
			clr: clientRepo,
			ctr: contactRepo,
		}

		client := &models.Client{ID: "client-123", Name: "Gabriel"}

		clientRepo.
			On("GetClientByID", ctx, "client-123").
			Return(client, nil)

		contactRepo.
			On("CreateContact", ctx, mock.Anything).
			Return(errors.New("create contact error"))

		result, err := service.CreateContact(ctx, "123456789", "test@example.com", "client-123")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "create contact")
	})
}
