build:
	@go build -o bin/go-crud cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-crud