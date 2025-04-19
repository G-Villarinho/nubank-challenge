package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/g-villarinho/nubank-challenge/mocks"
	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientHandler_CreateClient(t *testing.T) {
	e := echo.New()
	ctx := context.Background()

	t.Run("should create client successfully", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		payload := `{
			"name": "Gabriel",
			"contacts": [
				{
					"phone": "+5521999999999",
					"email": "gabriel@gmail.com",
					"clientId": "ignored"
				}
			]
		}`

		req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer([]byte(payload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		clientService.
			On("CreateClient", ctx, "Gabriel", mock.Anything).
			Return(&models.ClientResponse{
				ID:        "client-123",
				Name:      "Gabriel",
				CreatedAt: time.Now(),
				Contacts: []*models.ContactResponse{
					{ID: "contact-1", Phone: "+5521999999999", Email: "gabriel@gmail.com"},
				},
			}, nil)

		err := handler.CreateClient(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		clientService.AssertExpectations(t)
	})

	t.Run("should return 400 on bad JSON", func(t *testing.T) {
		handler := &clientHandler{}
		req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer([]byte(`invalid-json`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := handler.CreateClient(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 500 if service fails", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		payload := `{
			"name": "Gabriel",
			"contacts": [
				{
					"phone": "+5521999999999",
					"email": "gabriel@gmail.com",
					"clientId": "ignored"
				}
			]
		}`

		req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer([]byte(payload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		clientService.
			On("CreateClient", ctx, "Gabriel", mock.Anything).
			Return(nil, errors.New("internal error"))

		err := handler.CreateClient(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestClientHandler_GetClientsWithContact(t *testing.T) {
	e := echo.New()
	ctx := context.Background()

	t.Run("should return clients with contacts", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		mockResponse := []models.ClientResponse{
			{
				ID:        "client-123",
				Name:      "Gabriel",
				CreatedAt: time.Now(),
				Contacts: []*models.ContactResponse{
					{
						ID:    "contact-1",
						Phone: "+5521999999999",
						Email: "gabriel@gmail.com",
					},
				},
			},
		}

		clientService.
			On("GetClientsWithContact", ctx).
			Return(mockResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/clients", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		err := handler.GetClientsWithContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		clientService.AssertExpectations(t)
	})

	t.Run("should return 500 if service fails", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		clientService.
			On("GetClientsWithContact", ctx).
			Return(nil, errors.New("unexpected failure"))

		req := httptest.NewRequest(http.MethodGet, "/clients", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		err := handler.GetClientsWithContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestClientHandler_GetClientContactsByID(t *testing.T) {
	e := echo.New()
	ctx := context.Background()

	t.Run("should return contacts by client ID", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		mockContacts := []models.ContactResponse{
			{
				ID:    "contact-1",
				Phone: "+5521999999999",
				Email: "gabriel@gmail.com",
			},
		}

		clientService.
			On("GetClientContactsByID", ctx, "client-123").
			Return(mockContacts, nil)

		req := httptest.NewRequest(http.MethodGet, "/clients/client-123/contacts", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("clientId")
		c.SetParamValues("client-123")
		c.SetRequest(req.WithContext(ctx))

		err := handler.GetClientContactsByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		clientService.AssertExpectations(t)
	})

	t.Run("should return 400 if clientId param is missing", func(t *testing.T) {
		handler := &clientHandler{}

		req := httptest.NewRequest(http.MethodGet, "/clients//contacts", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("clientId")
		c.SetParamValues("")

		err := handler.GetClientContactsByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 404 if client not found", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		clientService.
			On("GetClientContactsByID", ctx, "missing-client").
			Return(nil, models.ErrClientNotFound)

		req := httptest.NewRequest(http.MethodGet, "/clients/missing-client/contacts", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("clientId")
		c.SetParamValues("missing-client")
		c.SetRequest(req.WithContext(ctx))

		err := handler.GetClientContactsByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("should return 500 if service fails", func(t *testing.T) {
		clientService := new(mocks.ClientServiceMock)
		handler := &clientHandler{cs: clientService}

		clientService.
			On("GetClientContactsByID", ctx, "client-123").
			Return(nil, errors.New("unexpected failure"))

		req := httptest.NewRequest(http.MethodGet, "/clients/client-123/contacts", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("clientId")
		c.SetParamValues("client-123")
		c.SetRequest(req.WithContext(ctx))

		err := handler.GetClientContactsByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
