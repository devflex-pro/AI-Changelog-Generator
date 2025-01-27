BINARY_NAME := clitool
MAIN_DIR := cmd/cli-tool

build:
	go build -o $(BINARY_NAME) $(MAIN_DIR)/main.go
	chmod +x $(BINARY_NAME)
	@echo "Build complete: ./$(BINARY_NAME)"

run: build
	./$(BINARY_NAME) $(ARGS)

clean:
	rm -f $(BINARY_NAME)
	@echo "Clean complete"

lint:
	golangci-lint run ./...
	@echo "Linting complete"

test:
	go test ./... -v
	@echo "Tests complete"

check: lint test
	@echo "Code check complete"

help:
	@echo "Usage:"
	@echo "  make build     - Build the CLI tool"
	@echo "  make run ARGS='--repo https://github.com/user/repo' - Run the CLI tool with arguments"
	@echo "  make clean     - Remove the built binary"
	@echo "  make lint      - Run golangci-lint"
	@echo "  make test      - Run tests"
	@echo "  make check     - Run lint and tests"
