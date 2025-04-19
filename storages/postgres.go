package storages

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/nubank-challenge/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage(ctx context.Context) (*gorm.DB, error) {
	dsn := getDSN()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	slqDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	slqDB.SetMaxOpenConns(configs.Env.Postgres.MaxConn)
	slqDB.SetMaxIdleConns(configs.Env.Postgres.MaxIdle)
	slqDB.SetConnMaxLifetime(time.Duration(configs.Env.Postgres.MaxLifeTime) * time.Second)

	if err := slqDB.PingContext(ctx); err != nil {
		_ = slqDB.Close()
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		configs.Env.Postgres.Host,
		configs.Env.Postgres.Port,
		configs.Env.Postgres.User,
		configs.Env.Postgres.DBName,
		configs.Env.Postgres.Password,
		configs.Env.Postgres.DBSSLMode,
	)
}
