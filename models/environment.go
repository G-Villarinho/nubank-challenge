package models

type Environment struct {
	Env      string `env:"ENV"`
	Postgres Postgres
}

type Postgres struct {
	Host        string `env:"POSTGRES_HOST"`
	Port        int    `env:"POSTGRES_PORT"`
	User        string `env:"POSTGRES_USER"`
	Password    string `env:"POSTGRES_PASSWORD"`
	DBName      string `env:"POSTGRES_NAME"`
	DBSSLMode   string `env:"POSTGRES_SSL_MODE"`
	MaxConn     int    `env:"POSTGRES_MAX_CONN"`
	MaxIdle     int    `env:"POSTGRES_MAX_IDLE"`
	MaxLifeTime int    `env:"POSTGRES_MAX_LIFE_TIME"`
	Timeout     int    `env:"POSTGRES_TIMEOUT"`
}
