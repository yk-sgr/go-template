.PHONY: dev
dev:
	@echo "Starting development server..."
	@air

.PHONY: build
build:
	@echo "Building binary..."
	@go build -o ./bin/api ./cmd/api/main.go
	@echo "Done!"

.PHONY: new-migration
new-migration:
	@echo "Creating new migration with name $(name)..."
	@dbmate new $(name)

.PHONY: db-up
db-up:
	@echo "Running migrations..."
	@dbmate up

.PHONY: db-down
db-down:
	@echo "Rolling back migrations..."
	@dbmate down

.PHONY: db-gen
db-gen:
	@echo "Generating sqlc code..."
	@sqlc generate

.PHONY: swaggo-gen
swaggo-gen:
	@echo "Generating swagger docs..."
	@swag init -g internal/api/*.go

.PHONY: swaggo-fmt
swaggo-fmt:
	@echo "Formatting swagger docs..."
	@swag fmt -d internal/api