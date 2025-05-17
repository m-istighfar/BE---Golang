# Makefile

ENV_FILE := .env
include $(ENV_FILE)

MIGRATIONS_DIR := ./db/migrations
MIGRATE := migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)"

.PHONY: migrate-up migrate-down create force version clean

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

force:
	@read -p "Force version: " v; \
	$(MIGRATE) force $$v

version:
	$(MIGRATE) version

clean:
	rm -f $(MIGRATIONS_DIR)/*.sql

reset:
	$(MIGRATE) drop -f

schema:
	pg_dump --schema-only --no-owner --no-privileges -d "$(DB_URL)" > db/schema.sql

run:
	export APP_ENVIRONMENT=development && go run ./cmd/api/


.PHONY: generate-mocks

generate-mocks:
	mockery --name ProductUsecase --dir internal/usecase --output testing/mocks
	mockery --name ProductRepository --dir internal/repository --output testing/mocks

.PHONY: test test-unit test-all test-product-handler test-product-usecase test-product-repository

test-all: test-unit test-product-handler test-product-usecase test-product-repository
	@echo "All tests completed"

test: test-all

test-unit:
	go test -v ./testing/unit/...

test-product-handler:
	go test -v ./testing/unit/product_handler_test.go

test-product-usecase:
	go test -v ./testing/unit/product_usecase_test.go

test-product-repository:
	go test -v ./testing/unit/product_repository_test.go