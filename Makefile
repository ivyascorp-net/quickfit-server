# Makefile for QuickFit Server

.PHONY: run seed migrate fmt lint test clean

run:
	go run cmd/main.go

seed:
	go run seed/seed.go

migrate:
	go run seed/migrate.go

fmt:
	gofmt -w .

lint:
	golangci-lint run || true

test:
	go test ./...

clean:
	rm -rf dist

kill:
	pkill -f 'go run cmd/main.go' || true
