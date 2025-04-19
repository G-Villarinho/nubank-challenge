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

// CreateClient godoc
// @Summary Cria um novo cliente com contatos
// @Tags clients
// @Accept json
// @Produce json
// @Param payload body models.CreateClientPayload true "Dados do cliente"
// @Success 201 {object} models.ClientResponse
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Router /clients [post]
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

// GetClientsWithContact godoc
// @Summary Lista todos os clientes com seus contatos
// @Description Retorna uma lista de clientes com os respectivos contatos associados
// @Tags clients
// @Produce json
// @Success 200 {array} models.ClientResponse
// @Failure 500 {object} nil "Erro interno ao buscar clientes"
// @Router /clients [get]
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

// GetClientContactsByID godoc
// @Summary Lista contatos de um cliente específico
// @Description Retorna os contatos associados a um cliente pelo ID
// @Tags clients
// @Produce json
// @Param clientId path string true "ID do cliente"
// @Success 200 {array} models.ContactResponse
// @Failure 400 {object} nil "ID inválido ou ausente"
// @Failure 404 {object} nil "Cliente não encontrado"
// @Failure 500 {object} nil "Erro interno ao buscar contatos"
// @Router /clients/{clientId}/contacts [get]
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
