DOCKERCOMPOSECMD=docker-compose

.PHONY: up down restart run pgquery

up:
	$(DOCKERCOMPOSECMD) up -d --force-recreate
	@echo "Waiting until Postgres be ready..."
	@until docker ps | grep db | grep "(healthy)"; do sleep 1; done
	@echo "Postgres is started."

down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

restart: down up

run:
	go run ./cmd/main.go ./cmd/container.go

pgquery:
	docker exec -it db psql -U root -h localhost -d students -p 5432