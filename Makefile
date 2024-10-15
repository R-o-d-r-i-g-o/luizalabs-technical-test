
# VARIABLES
GO=go
PKG=$(shell go list ./... | grep -v /mock)
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

.PHONY: dev-up
dev-up:
	@docker-compose -f ./infra/docker/docker-compose.yml up -d

.PHONY: dev-stop
dev-stop:
	@docker-compose -f ./infra/docker/docker-compose.yml stop

.PHONY: dev-down
dev-down:
	@docker-compose -f ./infra/docker/docker-compose.yml down

.PHONY: build
build:
	@echo "Building the application..."
	$(GO) build -o bin/app $(MAIN)

.PHONY: test
test:
	@echo "Running tests with coverage..."
	$(GO) test -count=1 $(PKG) -cover -coverprofile=$(COVERAGE_OUT) $(EXCLUDE_MOCKS)
	$(GO) tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(COVERAGE_OUT) $(COVERAGE_HTML)

.PHONY: install-swagger-cli
install-swagger-cli:
	@echo "Running install swagger..."
	$(GO) install github.com/swaggo/swag/cmd/swag@v1.16.3

.PHONY: refresh-swagger
refresh-swagger:
	@echo "Running swagger lint..."
	@swag fmt
	@echo "Running swagger docs..."
	@swag init -g cmd/main.go --parseDependency

.PHONY: install-mock-cli
install-mock-cli:
	@echo "Installing mock cli..."
	@go install github.com/golang/mock/mockgen

.PHONY: run-mock
run-mock:
	@echo "Creating mock files for zipcode use-case..."
	@mockgen -source="internal/features/zipcode/repository.go" -destination="internal/features/zipcode/mock/repository.go" -package="mock"
	@mockgen -source="internal/features/zipcode/service.go"    -destination="internal/features/zipcode/mock/service.go"    -package="mock"
	@mockgen -source="internal/features/zipcode/handler.go"    -destination="internal/features/zipcode/mock/handler.go"    -package="mock"

	@echo "Creating mock files for auth use-case..."
	@mockgen -source="internal/features/auth/repository.go" -destination="internal/features/auth/mock/repository.go" -package="mock"
	@mockgen -source="internal/features/auth/service.go"    -destination="internal/features/auth/mock/service.go"    -package="mock"
	@mockgen -source="internal/features/auth/handler.go"    -destination="internal/features/auth/mock/handler.go"    -package="mock"


	@echo "Creating mock files for swagger use-case..."
	@mockgen -source="internal/features/swagger/handler.go" -destination="internal/features/swagger/mock/handler.go"    -package="mock"

	@echo "Creating mock files for health use-case..."
	@mockgen -source="internal/features/health/handler.go" -destination="internal/features/health/mock/handler.go"     -package="mock"

	@echo "Creating mock files for middlewares internal package..."
	@mockgen -source="internal/pkg/middleware/token_middleware.go" -destination="internal/pkg/middleware/mock/token_middleware.go" -package="mock"
	@mockgen -source="internal/pkg/middleware/cache_middleware.go" -destination="internal/pkg/middleware/mock/cache_middleware.go" -package="mock"


.PHONY: run-kubernets
run-kubernets:
	@kubectl apply -f ./infra/k8s/

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  all                  - Install dependencies"
	@echo "  install              - Install Go dependencies"
	@echo "  run                  - Run the application"
	@echo "  build                - Build the application"
	@echo "  dev-up               - Start all resources used by app"
	@echo "  dev-down             - Delete all resources used by app"
	@echo "  dev-stop             - Stop all resources used by app"
	@echo "  test                 - Run tests with coverage"
	@echo "  clean                - Clean up build files"
	@echo "  help                 - Show this help message"
	@echo "  install-swagger-cli  - Install swagger cli globally"
	@echo "  refresh-swagger      - Refresh swagger docs"
	@echo "  run-kubernets        - Deploy kubernets infraestructure"
	@echo "  install-mock-cli     - Install mockgen cli globally"
	@echo "  run-mock             - Generate/Upgrade mock files automatically"
	@echo "\nall install run build test clean help"


