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
	docker build -t tally .

.PHONY: docker-run
docker-run:
	@echo "\033[0;32m» Running bookclub backend container\033[0;39m"
	docker run -d --name tally_api --env-file config.env  --network tally-compose -p 8081:8081 tally:latest

