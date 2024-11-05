ifneq (,$(wildcard ./.env))
    include .env
    export
endif

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	@tailwindcss -i internal/infra/view/css/app.css -o public/styles.css --watch 

templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	npx tailwindcss -i internal/infra/view/css/app.css -o public/styles.css
	@templ generate view
	@go build -tags dev -o bin cmd/main.go


migrate_up:
	goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB_NAME)?sslmode=disable" up

migrate_down:
	goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB_NAME)?sslmode=disable" down

migrate_down_all:
	goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB_NAME)?sslmode=disable" down-to 0

migrate_reset:
	goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DB_NAME)?sslmode=disable" reset

seed:
	@go run cmd/seed/main.go
