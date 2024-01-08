DSN="host=localhost port=5432 user=root password=root dbname=capstonedb sslmode=disable timezone=UTC connect_timeout=5"
PORT=8080
SECRET=6SbKzhTqLzEnMxgcldutBV3PAMY/7EKUMU1Jj+Bx
DB_DOCKER_CONTAINER=capstone_db
BINARY_NAME=capstoneapi

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root capstonedb

stop_containers:
	@echo "Stopping other docker containers"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers..."; \
		 docker stop $$(docker ps -q); \
	else \
		echo "no active containers found..."; \
	fi

start_docker:
	docker start ${DB_DOCKER_CONTAINER}


create_migrations:
	sqlx migrate add -r $(name)

migrate_up:
	sqlx migrate run --database-url "postgres://root:root@localhost:5432/capstonedb?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://root:root@localhost:5432/capstonedb?sslmode=disable"

build:
	@echo "Building backend api binary"
	go build -o ${BINARY_NAME} server/*.go
	@echo "Binary built!"

run: build
	@echo "Start api"
	@env PORT=${PORT} DSN=${DSN} ./${BINARY_NAME} &
	@echo "api started!"

stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backend"

start: run

restart: stop start

postgres_restart: stop_containers start-docker
