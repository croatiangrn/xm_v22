DOCKER_COMPOSE = docker compose -f docker/docker-compose.yml --env-file .env

DOCKER = docker

build: ##Build environment
	$(DOCKER_COMPOSE) build --no-cache

start: ##Start containers in detached mode
	$(DOCKER_COMPOSE) up -d

stop: ##Stop containers
	$(DOCKER_COMPOSE) stop

restart: ##Restart running containers
	make stop && make start

tests: ##Run tests locally
	go test -v ./...