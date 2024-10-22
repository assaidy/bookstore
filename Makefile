include .env
export $(shell sed 's/=.*//' .env)

GOOSE_DRIVER = postgres
GOOSE_DBSTRING = "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable&search_path=$(DB_SCHEMA)"
GOOSE_MIGRATION_DIR = "internals/database/migrations" 
GOOSE_ENV = GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR)

all: build

run: build 
	@./bin/api-server

build:
	@go build -o ./bin/api-server ./cmd/main.go

clean:
	@rm -rf bin

up:
	$(GOOSE_ENV) goose up

down:
	$(GOOSE_ENV) goose down

reset:
	$(GOOSE_ENV) goose reset

migration:
	$(GOOSE_ENV) goose create -s $(name) sql

