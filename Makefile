CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

run:
	go run cmd/main.go

rungo:
	go run main.go

proto-gen:
	./scripts/gen-proto.sh

migrate_up:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/pentagol -verbose up

migrate_down:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/pentagol -verbose down

migrate_force:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/pentagol -verbose force 0

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-gen:
	swag init -g api/router.go -o api/docs