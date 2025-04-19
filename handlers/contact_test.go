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
)

func TestCreateContactHandler(t *testing.T) {
	e := echo.New()
	ctx := context.Background()

	t.Run("should create contact successfully", func(t *testing.T) {
		contactService := new(mocks.ContactServiceMock)

		handler := &contactHandler{
			cs: contactService,
		}

		payload := `{
			"phone": "+5521999999999",
			"email": "gabriel@gmail.com",
			"clientId": "client-123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewBuffer([]byte(payload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		contactService.
			On("CreateContact", ctx, "+5521999999999", "gabriel@gmail.com", "client-123").
			Return(&models.ContactResponse{
				ID:        "contact-1",
				Phone:     "+5521999999999",
				Email:     "gabriel@gmail.com",
				CreatedAt: time.Now(),
			}, nil)

		err := handler.CreateContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		contactService.AssertExpectations(t)
	})

	t.Run("should return 400 on bad payload", func(t *testing.T) {
		handler := &contactHandler{}

		req := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewBuffer([]byte(`invalid-json`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.CreateContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 404 if client not found", func(t *testing.T) {
		contactService := new(mocks.ContactServiceMock)

		handler := &contactHandler{
			cs: contactService,
		}

		payload := `{
			"phone": "+5521999999999",
			"email": "gabriel@gmail.com",
			"clientId": "missing-client"
		}`

		req := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewBuffer([]byte(payload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		contactService.
			On("CreateContact", ctx, "+5521999999999", "gabriel@gmail.com", "missing-client").
			Return(nil, models.ErrClientNotFound)

		err := handler.CreateContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("should return 500 if service fails", func(t *testing.T) {
		contactService := new(mocks.ContactServiceMock)

		handler := &contactHandler{
			cs: contactService,
		}

		payload := `{
			"phone": "+5521999999999",
			"email": "gabriel@gmail.com",
			"clientId": "client-123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewBuffer([]byte(payload)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req.WithContext(ctx))

		contactService.
			On("CreateContact", ctx, "+5521999999999", "gabriel@gmail.com", "client-123").
			Return(nil, errors.New("unexpected failure"))

		err := handler.CreateContact(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
