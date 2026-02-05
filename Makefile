.PHONY: clean start stop help: ## Prints help for targets with comments

help:
	@echo "Cluster Options: "
	@echo "	start				Build + Run Cluster Detached Mode"
	@echo "	stop				Deletes Volumes + Stops Cluster"
	@echo " clean				Prunes Extra Volumes"

start:
	@echo "Starting Docker Compose..."
	@docker compose up --build -d

stop:
	@echo "Stopping Docker Compose..."
	@docker compose down -v

clean:
	@echo "Cleaning Volumes And Containers"
	@docker container prune -f
	@docker volume prune -f
