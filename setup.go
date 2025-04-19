package main

import (
	"context"
	"log"

	"github.com/g-villarinho/nubank-challenge/handlers"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/g-villarinho/nubank-challenge/repositories"
	"github.com/g-villarinho/nubank-challenge/services"
	"github.com/g-villarinho/nubank-challenge/storages"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initDependencies(ctx context.Context, di *pkgs.Di) {
	db, err := storages.NewPostgresStorage(ctx)
	if err != nil {
		log.Fatal(err)
	}

	pkgs.Provide(di, func(di *pkgs.Di) (*gorm.DB, error) {
		return db, nil
	})

	// Handlers
	pkgs.Provide(di, handlers.NewClientHandler)
	pkgs.Provide(di, handlers.NewContactHandler)

	// Services
	pkgs.Provide(di, services.NewClientService)
	pkgs.Provide(di, services.NewContactService)

	// Repositories
	pkgs.Provide(di, repositories.NewClientRepository)
	pkgs.Provide(di, repositories.NewContactRepository)
}

func setupRoutes(e *echo.Echo, di *pkgs.Di) {
	setupClientRoutes(e, di)
	setupContactRoutes(e, di)
}

func setupClientRoutes(e *echo.Echo, di *pkgs.Di) {
	clientHandler, err := pkgs.Invoke[handlers.ClientHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/clients", clientHandler.CreateClient)
	e.GET("/clients", clientHandler.GetClientsWithContact)
	e.GET("/clients/:clientId/contacts", clientHandler.GetClientContactsByID)
}

func setupContactRoutes(e *echo.Echo, di *pkgs.Di) {
	contactHandler, err := pkgs.Invoke[handlers.ContactHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/contacts", contactHandler.CreateContact)
}
