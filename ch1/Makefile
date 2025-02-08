.PHONY: fmt vet build run
.DEFAULT_GOAL := build

fmt:
	go fmt ./

vet:
	go vet ./

build: fmt vet
	go build

clean:
	rm hello_world

run: build
	./hello_world
