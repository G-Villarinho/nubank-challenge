.PHONY: docker-run
docker-run:
	@if docker compose up --build 2>/dev/null; then \
			: ; \
	else \
		@echo "Falling back to Docker compose V1"; \
		@docker-compose up --build; \
	fi

.PHONY: docker-down
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		@echo "Falling back to Docker Compose V1"; \
		@docker-compose down; \
	fi


.PHONY: run
run:
	@echo "Running application... \n"
	@go run main.go setup.go

.PHONY: migrations
migrations:
	@echo "Running migrations... \n"
	@go run migrations/migrations.go

.PHONY: swag
swag:
	@echo "Generating Swagger documentation... \n"
	@swag init --output ./docs
	@echo "\n Swagger documentation generated successfully! \n"

.PHONY: mock
mock: 
	@echo "🔄 Gen mock..."
	@mockery

.PHONY: test
test:
	@echo "Running tests... \n"
	@go test -v ./... -coverprofile=coverage.out
	@echo "\n Tests completed successfully! \n"
	@go tool cover -html=coverage.out -o coverage.html
	@echo "\n Coverage report generated successfully! \n"