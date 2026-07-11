# Makefile for building and running the Go microservice
.PHONY: backend
backend:
	@echo "Building and starting the Go microservice..."
	docker compose up --build
ิbackend-down:
	@echo "Stopping and removing the Go microservice..."
	docker compose down