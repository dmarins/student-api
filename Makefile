DOCKERCOMPOSECMD=docker-compose
GOCMD=go
DOCKERCMD=docker
GOBIN=$(shell $(GOCMD) env GOPATH)/bin

.PHONY: down local-up local-restart docker-up docker-restart download run pgquery mockgen-download mocks-clean mocks-gen tests-clean tests tests-coverage

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

download:
	$(GOCMD) mod download

run: download
	APP_ENV=local $(GOCMD) run ./cmd/main.go

pgquery:
	$(DOCKERCMD) exec -it db psql -U root -h localhost -d students -p 5432

mockgen-download: download
	$(GOCMD) install -mod=mod go.uber.org/mock/mockgen@latest

mocks-clean:
	@echo "CLEANING MOCKS START..."
	rm -rf internal/domain/mocks/*
	@echo "CLEAN MOCKS!"

mocks-gen: mockgen-download mocks-clean
	$(GOBIN)/mockgen -source=internal/infrastructure/logger/logger.go -destination=internal/domain/mocks/logger.go -typed=true -package=mocks	
	$(GOBIN)/mockgen -source=internal/infrastructure/tracer/tracer.go -destination=internal/domain/mocks/tracer.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/infrastructure/tracer/span_wrapper.go -destination=internal/domain/mocks/span_wrapper.go -typed=true -package=mocks 
	$(GOBIN)/mockgen -source=internal/domain/repositories/healthcheck_repository.go -destination=internal/domain/mocks/healthcheck_repository.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/domain/repositories/student_repository.go -destination=internal/domain/mocks/student_repository.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/domain/usecases/healthcheck.go -destination=internal/domain/mocks/healthcheck.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/domain/usecases/student_create.go -destination=internal/domain/mocks/student_create.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/domain/usecases/student_read.go -destination=internal/domain/mocks/student_read.go -typed=true -package=mocks
	$(GOBIN)/mockgen -source=internal/domain/usecases/student_update.go -destination=internal/domain/mocks/student_update.go -typed=true -package=mocks

tests-clean:
	$(GOCMD) clean -testcache

tests: tests-clean local-restart
	$(GOCMD) test -cover -p=1 ./...

tests-coverage: tests-clean
	$(GOCMD) test -cover -p=1 -covermode=count -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out