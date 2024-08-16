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
	up down

test:
	go test ./...

swagger:
	swag init