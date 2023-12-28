package config

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BindAddr    string   `envconfig:"BIND_ADDR" default:":8080"`
	Host        string   `envconfig:"HOST" default:"http://localhost:8080"`
	DatabaseURL string   `envconfig:"DATABASE_URL" default:"postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"`
	CorsOrigins []string `envconfig:"CORS_ORIGINS" default:"http://localhost:3000"`
	JWTSecret   string   `envconfig:"JWT_SECRET" default:"token"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		slog.Error("Failed to load config", err)
	}
	return &cfg
}
