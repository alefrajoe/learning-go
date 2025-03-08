.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o main main.go

.PHONY: tidy
tidy:
	@go mod tidy
	@go mod vendor

.PHONY: test
test:
	@go test -v ./...
