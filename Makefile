.PHONY: build test clean

build:
	go build ./cmd/main.go

test:
	go test ./...

clean:
	go clean
