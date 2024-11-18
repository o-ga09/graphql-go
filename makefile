.PHONY build:
build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

.PHONY run:
run:
	@echo "Running..."
	@go run cmd/$(APP_NAME)/main.go

.PHONY test:
test:
	@echo "Testing..."
	@go test -v ./...

.PHONY clean:
clean:
	@echo "Cleaning..."
	@rm -rf db/db
	@rm -rf graph
	@rm -f server.go
	@rm -rf bin/$(APP_NAME)

.PHONY generate:
generate:
	@echo "Generating..."
	@sqlc -f internal/db/sqlc.yml generate
	@gqlgen generate
	@go generate ./...
