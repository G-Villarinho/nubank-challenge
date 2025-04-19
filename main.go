// @title Nubank Challenge API
// @version 1.0
// @description Esta API gerencia clientes e contatos.
// @host localhost:8080
// @BasePath /
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/nubank-challenge/configs"
	_ "github.com/g-villarinho/nubank-challenge/docs"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	if err := configs.LoadEnv(); err != nil {
		e.Logger.Fatal(fmt.Sprintf("load env: %v", err))
	}

	di := pkgs.NewDi()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Use(middleware.Recover())

	initDependencies(ctx, di)
	setupRoutes(e, di)

	if configs.Env.Env == "DEV" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
