.PHONY: start pgsql-verify pgsql stop clean sqlc-generate 

GO_RUN := go run ./cmd/api-server/main.go 

run: go.mod pgsql-verify
	@$(GO_RUN) &

build:
	@go build -o ./bin/exec ./cmd/api-server

start: pgsql-verify
	@go run ./bin/exec

# execute in container pg, the cli, with username, on database, this command
pgsql-verify: pgsql
	@echo "Verifying Postgres Container Network..."
	@docker exec pg psql -U postgres -d localDB -c "SELECT 1;"
	@echo "Postgres Container Verified"

# show all, without all the extra info, filter by name, else run container with env in the background on this image
pgsql:
	@echo "Starting Docker Container..."
	@if [ $$(docker ps -a -q -f name=pg) ]; then \
		docker start pg; \
	else \
		docker run --name pg \
			-e POSTGRES_USER=postgres \
			-e POSTGRES_PASSWORD=postgres \
			-e POSTGRES_DB=localDB \
			-p 5432:5432 \
			-d postgres; \
	fi
	@sleep 3

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
