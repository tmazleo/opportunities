.PHONY: default run buld test docs clean

APP_NAME=gopportunities

default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@$(shell go env GOPATH)/bin/swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@$(shell go env GOPATH)/bin/swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs