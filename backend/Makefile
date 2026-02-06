# this file is for just server testing, THERE IS NO DATABASE
.PHONY: start stop clean sqlc-generate build

GO_RUN := go run ./cmd/api-server/main.go 

run: go.mod pgsql-verify
	@$(GO_RUN) &

build:
	@go build -o server ./cmd/api-server

start:
	@go run server

# Finds Process and Kills the Server with the Postgres Container
stop:
	@echo "Stopping Server..."
	@pkill -u $$(whoami) -f $(GO_RUN)

clean:
	@docker container prune -f
	@docker volume prune -f

sqlc-generate:
	@sqlc generate

install-tools:
	@go get "github.com/jackc/pgx/v5"
	@go get "github.com/sqlc-dev/sqlc/cmd/sqlc@latest"

