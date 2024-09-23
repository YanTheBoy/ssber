package config

import (
	"os"
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
	ServerAddress    string        `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8082"`
	Timeout          time.Duration `env:"TIMEOUT" env-default:"5s"`
}


func Parse() (Config, error) {
	cfg := Config{
		ServerAddress:    getenv("SERVER_ADDRESS", "0.0.0.0:8082"),
		PostgresConn:     getenv("POSTGRES_CONN", ""),
		PostgresJdbcURL:  getenv("POSTGRES_JDBC_URL", ""),
		PostgresUsername: getenv("POSTGRES_USERNAME", ""),
		PostgresPassword: getenv("POSTGRES_PASSWORD", ""),
		PostgresHost:     getenv("POSTGRES_HOST", ""),
		PostgresPort:     getenv("POSTGRES_PORT", "5432"),
		PostgresDatabase: getenv("POSTGRES_DATABASE", ""),
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}