BIN_NAME=bookclub-backend

wire:
	@echo building wire....
	@wire

build:
	@echo building binary...
	@GOPRIVATE=github.com/epociask CGO_ENABLED=0 go build -a -tags netgo -o bin/$(BIN_NAME);

run: 
	@./bin/${BIN_NAME}

test:
	@ go test ./...

.PHONY: build-linux
build-linux:
	@echo "\033[0;32m» Building bookclub backend binary \033[0;39m"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o bin/$(APP)

.PHONY: docker-build
docker-build:  build-linux
	@echo "\033[0;32m» Building bookclub backend image \033[0;39m"
	docker build -t bookclub_api .

.PHONY: docker-run
docker-run:
	@echo "\033[0;32m» Running bookclub backend container\033[0;39m"
	docker run -d --name bookclub-api --env-file config.env  --network bookclub-compose -p 8081:8081 bookclub_api:latest

.PHONY: docker-up
docker-up:
	@echo "\033[0;32m» Building bookclub backend dependencies\033[0;39m"
	docker-compose up -d

