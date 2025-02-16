build:
	@go build -o bin/go-crud cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-crud

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrations/main.go up

migrate-down:
	@go run cmd/migrate/migrations/main.go down