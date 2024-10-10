
# VARIABLES
GO=go
PKG=./...
MAIN=./cmd/main.go
COVERAGE_OUT=coverage.out
COVERAGE_HTML=coverage.html

# DEFAULT TARGET
all: install

.PHONY: install
install:
	@echo "Installing dependencies..."
	$(GO) mod tidy
	$(GO) mod vendor

.PHONY: run
run:
	@echo "Starting app..."
	$(GO) run $(MAIN)

.PHONY: build
build:
	@echo "Building the application..."
	$(GO) build -o bin/app $(MAIN)

.PHONY: test
test:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=$(COVERAGE_OUT) $(PKG)
	$(GO) tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  all       - Install dependencies"
	@echo "  install   - Install Go dependencies"
	@echo "  run       - Run the application"
	@echo "  build     - Build the application"
	@echo "  test      - Run tests with coverage"
	@echo "  clean     - Clean up build files"
	@echo "  help      - Show this help message"
	@echo "\nall install run build test clean help"


