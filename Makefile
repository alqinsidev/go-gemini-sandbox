.PHONY: migrate
migrate:
	@go run src/db/migration/main.go

.PHONY: schema
schema:
	@migrate create -ext sql -dir src/db/migration/scripts -seq $(name)

.PHONY: lint
lint:
	@golangci-lint run -v

.PHONY: validate
validate:
	@go test -v ./...
	@golangci-lint run -v
	@go build -v ./...
	