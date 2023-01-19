SHELL=cmd.exe

SEARCH_BINARY=Dockerfile



## up: starts all containers in the background without forcing build

build:

	@echo Starting Docker images...

	docker-compose build

	@echo Docker images started!



## up_build: stops docker-compose (if running), builds all projects and starts docker compose

up:

	@echo Stopping docker images (if running...)

	docker-compose down --rmi local

	@echo Building (when required) and starting docker images...

	docker-compose up 

	@echo Docker images built and started!



## down: stop docker compose

down:

	@echo Stopping docker compose...

	docker-compose down

	@echo Done!