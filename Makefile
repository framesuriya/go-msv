# Makefile for building and running the Go microservice
.PHONY: backend backend-down
backend:
	@echo "Building and starting the Go microservice..."
	docker compose up --build
ิbackend-down:
	@echo "Stopping and removing the Go microservice..."
	docker compose down
backend-down-volumes:
	@echo "Stopping and removing the Go microservice and its volumes..."
	docker compose down -v