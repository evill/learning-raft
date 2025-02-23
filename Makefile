.PHONY: build run test clean fmt lint

# Variables
BINARY_NAME=raft-server
GO=go
HTTP_NATIVE_GENERATED_TRANSPORT_PATH=internal/transports/http/inbound/native

# Build the application
build: generate
	$(GO) build -o ./bin/$(BINARY_NAME) ./cmd/service/main.go

# Run the application
run: build
	exec ./bin/$(BINARY_NAME)

# Run the application
rundev: generate
	exec env RAFTIG_NODE_INTERNAL_URI=http://localhost:8081 \
		RAFTIG_API_INTERNAL_PORT=8081 \
		$(GO) run ./cmd/service/main.go

# Run tests
test:
	$(GO) test ./tests/integration

# Clean build artifacts
clean: cleanapi
	rm -rf ./bin/**
	$(GO) clean
# Format code
fmt:
	$(GO) fmt ./...

# Run linter
lint:
	golangci-lint run ./...

# Install dependencies
deps: generate
	$(GO) mod tidy
	$(GO) mod verify

generate: genapi

genapi: cleanapi
	./tools/generate/openapi/install-openapi-generator.sh
	./tools/generate/openapi/generate-openapi-native.sh

cleanapi:
	rm -rf $(HTTP_NATIVE_GENERATED_TRANSPORT_PATH)
