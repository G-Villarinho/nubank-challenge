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