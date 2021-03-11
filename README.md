# Bookclub-Backend 

## Setup
	1. `cp config.env.template config.env`
	2. Ensure you have golang installed 
	3. Ensure you have wire installed... if not run `go get github.com/google/wire/cmd/wire`

## To Run Binary 
	1. Make sure you have wire installed 
	2. Run `make build` to build binary
	3. Run  `make run` to execute the binary 

## To Develop With Docker 
	1. Run `make docker-up` to spin up API network dependencies 
	2. Run `make docker-build` to compile go module into a linux binary and embed it as a docker image
	3. Run `make docker-run` to run the service

## To Test it's properly working 
	1. Run `curl --location --request GET 'http://localhost:{PORT}/health' `, it should return a json health model with a current timestamp and boolean set to true 

## Running Tests 
* For unit tests, run `make test` to run all associated unit tests included in `/internal`

* For integration tests:
	1. Ensure the API and its dependencies are running 
	2. Run `make integration-tests` to run a test harness against your local docker network

## Notes
* You can reference the current port from the `config.env.template` or your `config.env`
* If you augment the parameters/return-vals of the wire binded structs (any package with `var Module` in it), ensure to run `make wire` to correctly update these code changes in your `wire_gen.go` file
* If you're getting an error describing that `interface does not implement {STRUCT_NAME}`, make sure that the prototypes of all receiver functions that are publicly defined with the package are embedded within the interface  

