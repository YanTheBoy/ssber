package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME"`
	PostgresConn     string        `env:"POSTGRES_CONN"`
	PostgresJdbcURL  string        `env:"POSTGRES_JDBC_URL"`
	PostgresUsername string        `env:"POSTGRES_USERNAME"`
	PostgresPassword string        `env:"POSTGRES_PASSWORD"`
	PostgresHost     string        `env:"POSTGRES_HOST"`
	PostgresPort     string        `env:"POSTGRES_PORT"`
	PostgresDatabase string        `env:"POSTGRES_DATABASE"`
	ServerAddress    string        `env:"SERVER_ADDRESS"`
	Timeout          time.Duration `env:"TIMEOUT" env-default:"5s"`
}


func Parse() (Config, error) {
	// Сначала ищем .ENV файл, если нету, то берем ENV параметров
	err := godotenv.Load()
	if err != nil {
		log.Println("Не найден .ENV файл,")
	}
	var cfg Config

	err = env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

