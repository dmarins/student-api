DOCKERCOMPOSECMD=docker-compose

.PHONY: up down restart test

up:
	$(DOCKERCOMPOSECMD) up -d --force-recreate
	@echo "Waiting until Mysql be ready..."
	@until docker ps | grep mysql | grep "(healthy)"; do sleep 1; done
	@echo "Mysql is started."

down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

restart:
	down up

test:
	go test ./...

swagger:
	swag init

run:
	go run ./cmd/main.go ./cmd/container.go

db-init:
	docker exec -it mysql mysql -uroot -proot students create table students (id varchar(36) not null, name varchar(200) not null, primary key (id) );

db-query:
	docker exec -it mysql mysql -uroot -proot students