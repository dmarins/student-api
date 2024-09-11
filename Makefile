DOCKERCOMPOSECMD=docker-compose

.PHONY: up down restart test

up:
	$(DOCKERCOMPOSECMD) up -d --force-recreate
	@echo "Waiting until Postgres be ready..."
	@until docker ps | grep db | grep "(healthy)"; do sleep 1; done
	@echo "Postgres is started."

down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

restart:
	down up

run:
	go run ./cmd/main.go ./cmd/container.go

db-init:
	docker exec -it db psql -U root -h localhost -d students -p 5432 -c "CREATE TABLE students (id VARCHAR(36) NOT NULL, name VARCHAR(200) NOT NULL, PRIMARY KEY (id));"

db-query:
	docker exec -it db psql -U root -h localhost -d students -p 5432