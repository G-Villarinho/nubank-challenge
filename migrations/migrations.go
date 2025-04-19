package main

import (
	"context"
	"log"
	"time"

	"github.com/g-villarinho/nubank-challenge/configs"
	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/g-villarinho/nubank-challenge/storages"
)

func main() {
	configs.LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := storages.NewPostgresStorage(ctx)
	if err != nil {
		log.Fatal("connect to database: ", err)
	}

	err = db.AutoMigrate(
		&models.Client{},
		&models.Contact{},
	)

	if err != nil {
		log.Fatal("auto migrate: %w", err)
	}

	log.Println("migrations excuted succefully!")
}
