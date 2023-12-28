package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/yk-sgr/nexthook-api/internal/api"
	"github.com/yk-sgr/nexthook-api/internal/config"
	"github.com/yk-sgr/nexthook-api/internal/db"
	"github.com/yk-sgr/nexthook-api/internal/services"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	authService := services.NewAuthService(queries, cfg.JWTSecret)
	userService := services.NewUserService(queries)

	api := api.New(&api.Options{
		CorsAllowedOrigins: cfg.CorsOrigins,
		AuthService:        authService,
		UserService:        userService,
	})
	api.Start(cfg.BindAddr)
}
