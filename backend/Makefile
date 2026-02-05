.PHONY: start pgsql-verify stop clean sqlc-generate build

GO_RUN := go run ./cmd/api-server/main.go 

run: go.mod pgsql-verify
	@$(GO_RUN) &

build:
	@go build -o ./bin ./cmd/api-server

start:
	@go run ./bin

# execute in container pg, the cli, with username, on database, this command
pgsql-verify: 
	@echo "Verifying Postgres Container Network..."
	@docker exec pg psql -U postgres -d localDB -c "SELECT 1;"
	@echo "Postgres Container Verified"

sqlc-generate:
	@sqlc generate

install-tools:
	@go get "github.com/jackc/pgx/v5"
	@go get "github.com/sqlc-dev/sqlc/cmd/sqlc@latest"

# Finds Process and Kills the Server with the Postgres Container
stop:
	@echo "Stopping Server..."
	@pkill -u $$(whoami) -f $(GO_RUN)
	@echo "Stopping Postgres Container..."
	@docker stop pg

clean:
	@docker container prune -f
	@docker volume prune -f
