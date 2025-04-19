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

type ClientHandler interface {
	CreateClient(ectx echo.Context) error
	GetClientsWithContact(ectx echo.Context) error
	GetClientContactsByID(ectx echo.Context) error
}

type clientHandler struct {
	di *pkgs.Di
	cs services.ClientService
}

func NewClientHandler(di *pkgs.Di) (ClientHandler, error) {
	clientService, err := pkgs.Invoke[services.ClientService](di)
	if err != nil {
		return nil, fmt.Errorf("invoke services.client: %w", err)
	}

	return &clientHandler{
		di: di,
		cs: clientService,
	}, nil
}

func (c *clientHandler) CreateClient(ectx echo.Context) error {
	logger := slog.With(
		slog.String("handler", "client"),
		slog.String("method", "CreateClient"),
	)

	var payload models.CreateClientPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	contacts := models.ToContacts(payload.Contacts)

	response, err := c.cs.CreateClient(ectx.Request().Context(), payload.Name, contacts)
	if err != nil {
		logger.Error("error to create client", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusCreated, response)
}

func (c *clientHandler) GetClientsWithContact(ectx echo.Context) error {
	logger := slog.With(
		slog.String("handler", "client"),
		slog.String("method", "GetClientsWithContact"),
	)

	clients, err := c.cs.GetClientsWithContact(ectx.Request().Context())
	if err != nil {
		logger.Error("error to get clients with contact", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, clients)
}

func (c *clientHandler) GetClientContactsByID(ectx echo.Context) error {
	logger := slog.With(
		slog.String("handler", "client"),
		slog.String("method", "GetClientContactsByID"),
	)

	id := ectx.Param("clientId")
	if id == "" {
		return ectx.NoContent(http.StatusBadRequest)
	}

	response, err := c.cs.GetClientContactsByID(ectx.Request().Context(), id)
	if err != nil {
		if err == models.ErrClientNotFound {
			logger.Error("client not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("error to get client contacts by id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, response)
}
