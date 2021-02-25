# Go-REST-API-Template 

* Somewhat elegant repository to give HTTP REST API starter code in Go utilizing `github.com/google/wire` for dependency injection and building 

* Intended for fellow gophers and USFCA professors trying to learn better go practices (I'm talking to you Matt)

## Setup
	1. `cp config.env.template config.env`
	2. Ensure you have golang installed 
	3. Ensure you have wire installed... if not run `go get github.com/google/wire/cmd/wire`
## To run 
	1. Make sure you have wire installed 
	2. Run `make build` to build binary
	3. Run  `make run` to execute the binary 

## To test it's properly working 
	1. Run `curl --location --request GET 'http://localhost:8080/health' `, it should return a json health model with a current timestamp and boolean set to true 
