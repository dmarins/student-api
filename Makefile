DOCKERCOMPOSECMD=docker-compose
GOCMD=go
DOCKERCMD=docker

.PHONY: down local-up local-restart docker-up docker-restart run pgquery

down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

local-up:
	$(DOCKERCOMPOSECMD) -f docker-compose.local.yaml up -d
	@echo "Waiting until Postgres be ready..."
	@until docker ps | grep db | grep "(healthy)"; do sleep 1; done
	@echo "Postgres is started."

local-restart: down local-up

docker-up:
	$(DOCKERCOMPOSECMD) -f docker-compose.yaml up -d --build
	@echo "Waiting until Postgres be ready..."
	@until docker ps | grep db | grep "(healthy)"; do sleep 1; done
	@echo "Postgres is started."

docker-restart: down docker-up

run:
	APP_ENV=local $(GOCMD) run ./cmd/main.go

pgquery:
	$(DOCKERCMD) exec -it db psql -U root -h localhost -d students -p 5432