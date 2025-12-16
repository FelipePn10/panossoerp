include .env

PWD := $(shell pwd)

MIGRATIONS_DIR := $(PWD)/migrations

create_migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) -seq init

migrate_up:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		-verbose up

migrate_down:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		-verbose down

migrate_force:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		force 1

reset:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		-drop -verbose

.PHONY: create_migration migrate_up migrate_down migrate_force reset
