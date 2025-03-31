include cmd/api/.env

ifneq ($(filter migrate,$(MAKECMDGOALS)),)
  MIGRATION_NAME := $(word 2,$(MAKECMDGOALS))
endif

run:
	@echo "$(GREEN) Running Golang App... $(CYAN)LOCAL $(NC)"
	go run cmd/api/main.go

migrate:
	@echo "Creating migration: $(MIGRATION_NAME)"
	migrate create -ext sql -dir cmd/migrate -seq $(MIGRATION_NAME)
%:
	@:


sqlc:
	@echo "$(GREEN) Generated SQLC"
	sqlc generate
