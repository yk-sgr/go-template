package api

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yk-sgr/nexthook-api/docs"
	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

// Options is a set of options for the API.
type Options struct {
	CorsAllowedOrigins []string
	Host               string
	AuthService        domain.AuthService
	UserService        domain.UserService
}

// API is a the Nexthook REST API.
type API struct {
	r           chi.Router
	authService domain.AuthService
	userService domain.UserService
}

//	@title			Nexthook API
//	@version		1.0
//	@description	The Nexthook API.
//	@BasePath		/v1

// @SecurityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func New(opts *Options) *API {
	r := chi.NewRouter()
	api := &API{
		r:           r,
		authService: opts.AuthService,
		userService: opts.UserService,
	}

	// Register Middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   opts.CorsAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Swagger
	docs.SwaggerInfo.Host = opts.Host
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Register Routes
	r.Post("/v1/auth/signup", api.handleSignUp)
	r.Post("/v1/auth/signin", api.handleSignIn)
	r.Get("/v1/users/{id}", api.handleGetUserByID)

	return api
}

// Start starts the API server.
func (api *API) Start(addr string) {
	slog.Info("Starting server on " + addr)
	err := http.ListenAndServe(addr, api.r)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
