include .env
export

start:
	go build -o ./bin/app ./cmd/main.go
	./bin/app

database_up:
	docker run --name dev-db --rm \
	-e POSTGRES_USER=${DB_USER} \
	-e POSTGRES_PASSWORD=${DB_PASS} \
	-e POSTGRES_DB=${DB_NAME} \
	-e PGDATA=/var/lib/postgresql/data \
	-p 5432:5432 \
	-v ./migrations:/docker-entrypoint-initdb.d \
	-d postgres:15.3-bullseye

	clear

database_down:
	docker stop dev-db