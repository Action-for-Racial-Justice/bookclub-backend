BIN_NAME=bookclub-backend
D_NAME=bookclub_api

.PHONY: lint
lint:
	@echo "\033[0;32m» Linting Go code...\033[0;39m"
	@tempdir=$(mktemp -d);cd $(tempdir); GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21.0 2> /dev/null;rm -rf $(tempdir)
	@golangci-lint run


.PHONY: wire
wire:
	@echo building wire....
	@wire

.PHONY: lint build 
build:
	@echo building binary...
	@GOPRIVATE=github.com/epociask CGO_ENABLED=0 go build -a -tags netgo -o bin/$(BIN_NAME);

.PHONY: run 
run: 
	@./bin/${BIN_NAME}

.PHONY: test
test:
	@ go test ./... --cover

.PHONY: build-linux
build-linux:
	@echo "\033[0;32m» Building bookclub backend binary \033[0;39m"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o bin/$(BIN_NAME)

.PHONY: docker-build
docker-build:  build-linux
	@echo "\033[0;32m» Building bookclub backend image \033[0;39m"
	docker build -t $(BIN_NAME) .

.PHONY: docker-run
docker-run:
	@echo "\033[0;32m» Running bookclub backend container\033[0;39m"
	docker run -d --name $(D_NAME) --env-file config.env  --network bookclub-compose -p 8081:8081 $(BIN_NAME):latest

.PHONY: docker-up
docker-up:
	@echo "\033[0;32m» Building bookclub backend dependencies\033[0;39m"
	docker-compose up -d

.PHONY: gen-mocks
gen-mocks:
	@echo "\033[0;32m» Generating mocks... \033[0;39m"
	@GO111MODULE=on go generate --run "mockgen*" ./...


.PHONY: docker-api-down
docker-api-down:
	@docker stop $(D_NAME)
	@docker rm $(D_NAME)

.PHONY: docker-down
docker-down:
	docker-compose down 

.PHONY: integration-tests
integration-tests:
	@echo "\033[0;32m» Starting integration tests\033[0;39m"
	@echo "\033[0;32m» Building and running pytests \033[0;39m"
	@docker build --network bookclub-compose -t bookclub_pytest -f integration_tests/Dockerfile .
	@docker rmi bookclub_pytest

