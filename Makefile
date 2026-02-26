.PHONY: build run clean

BINARY_NAME=server
CMD_PATH=cmd/server/main.go

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) $(CMD_PATH)

run:
	@echo "Running $(BINARY_NAME)..."
	@go run $(CMD_PATH)

clean:
	@echo "Cleaning..."
	@rm -rf bin
