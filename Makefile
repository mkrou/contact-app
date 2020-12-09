include .env

# Project initialization
.PHONY: init
init:
	@cp dev.env .env
	@go mod download

# Docker
.PHONY: compose
compose:
	@docker-compose -f deployments/docker-compose.yml up -d

# Migrations
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable
MIGRATES_DB=-database "$(DB_URL)"
MIGRATES_DIR=scripts/postgres/migrations/
MIGRATES_DIR_ARG=-source file://$(MIGRATES_DIR)

.PHONY: migration
migration:
	@migrate $(MIGRATES_DB) create -ext sql -dir $(MIGRATES_DIR) -seq $(name)

.PHONY: migrate
migrate:
	@migrate $(MIGRATES_DIR_ARG) $(MIGRATES_DB) up

.PHONY: migrate-down
migrate-down:
	@migrate $(MIGRATES_DIR_ARG) $(MIGRATES_DB) down -all

.PHONY: remigrate
remigrate: migrate-down migrate

# app
.PHONY: run
run:
	@go run cmd/contact/main.go -c=.env