package main

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/nubank-challenge/configs"
	"github.com/g-villarinho/nubank-challenge/pkgs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Logger.Fatal(e.Start(":8080"))
}
