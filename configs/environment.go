package configs

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/g-villarinho/nubank-challenge/models"
	"github.com/joho/godotenv"
)

var Env models.Environment

func LoadEnv() error {
	if err := godotenv.Load(".env.local"); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	if _, err := env.UnmarshalFromEnviron(&Env); err != nil {
		return fmt.Errorf("init env: %w", err)
	}

	return nil
}
