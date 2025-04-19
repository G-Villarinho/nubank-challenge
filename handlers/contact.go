package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/g-villarinho/nubank-challenge/services"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type ContactHandler interface {
	CreateContact(ectx echo.Context) error
}

type contactHandler struct {
	di *pkgs.Di
	cs services.ContactService
}

func NewContactHandler(di *pkgs.Di) (ContactHandler, error) {
	cs, err := pkgs.Invoke[services.ContactService](di)
	if err != nil {
		return nil, fmt.Errorf("invoke services.contact: %w", err)
	}

	return &contactHandler{
		di: di,
		cs: cs,
	}, nil
}

// CreateContact godoc
// @Summary Cria um novo contato
// @Description Cria um novo contato associado a um cliente existente
// @Tags contacts
// @Accept json
// @Produce json
// @Param payload body models.CreateContactPayload true "Dados do contato"
// @Success 201 {object} models.ContactResponse
// @Failure 400 {object} nil "Erro de validação ou payload inválido"
// @Failure 404 {object} nil "Cliente não encontrado"
// @Failure 500 {object} nil "Erro interno ao criar contato"
// @Router /contacts [post]
func (c *contactHandler) CreateContact(ectx echo.Context) error {
	logger := slog.With(
		slog.String("handler", "contact"),
		slog.String("method", "CreateContact"),
	)

	var payload models.CreateContactPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("decode payload", slog.Any("error", err))
		return ectx.NoContent(http.StatusBadRequest)
	}

	response, err := c.cs.CreateContact(ectx.Request().Context(), payload.Phone, payload.Email, payload.ClientID)
	if err != nil {
		logger.Error("create contact", slog.Any("error", err))
		if err == models.ErrClientNotFound {
			return ectx.NoContent(http.StatusNotFound)
		}
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusCreated, response)
}
