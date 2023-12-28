# Go Template

## Installlation
### dbmate
#### NPM
`npm install --save-dev dbmate`
#### MacOS
`brew install dbmate`

### SQLC
`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### Swaggo
`go install github.com/swaggo/swag/cmd/swag@latest`

## Environment Variables
- `BIND_ADDR`: The address to bind the HTTP server to.
- `HOST`: THe API host address.
- `DATABASE_URL`: The Postgres DSN.
- `CORS_ORIGINS`: CORS origins as comma-separated value.
- `JWT_SECRET`: The secret for signing JWT tokens.

## Commands
- `make dev`: Starts a dev server with hot reload.
- `make build`: Builds the API binary.
- `make new-migration`: Create a new migration.
- `make db-up`: Applies database migrations.
- `make db-down`: Rolls back database migrations.
- `make db-gen`: Generates SQLC code.
- `make swaggo-gen`: Generates Swagger documentation.
- `make swaggo-fmt`: Formats Swaggo documentation.
  
## Project Structure
- `./cmd`: This directory contains the main applications of the project. Each subdirectory here represents a different executable.
  - `./cmd/api`: Houses the main entry point for the API server.

- `./internal`: This is where the core logic of the application resides. It's not accessible from other projects.
  - `./internal/config`: Contains configuration-related code.
  - `./internal/db`: Holds database-related functionalities.
  - `./internal/api`: Includes handlers and logic specific to the API.
  - `./internal/services`: Comprises the business logic and service layer.

- `./pkg`: Contains library code that's ok to use by external applications.
  - `./pkg/domain`: Includes domain-specific logic and entities.

- `./docs`: Swagger documentation files for the API.

- `./db`: Database related files.
  - `./db/migrations`: Database migration scripts.
  - `./db/query.sql`: SQLC query.sql file.
  - `./db/schema.sql`: Auto-generated database schema file by dbmate.