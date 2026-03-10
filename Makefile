build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test ./...

lint:
	golangci-lint run