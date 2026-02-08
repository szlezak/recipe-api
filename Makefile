# Variables
BINARY_NAME=recipe-api

## build: Build the docker image
build:
	docker-compose build

## up: Start the application in docker
up:
	docker-compose up

## down: Stop the application
down:
	docker-compose down

## restart: Restart the container and rebuild (Useful for code changes)
restart:
	docker-compose down
	docker-compose up --build -d

## logs: View logs from the container
logs:
	docker-compose logs -f

## clean: Remove the binary and the database file
clean:
	rm -f $(BINARY_NAME)
	rm -f recipes.db