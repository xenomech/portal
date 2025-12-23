.PHONY: build clean run

BINARY_NAME=portal
BUILD_DIR=./dist

VERSION=0.0.1
VERSION_NAME=alpha

LDFLAGS=-ldflags "\
	-X portal/internal/cli.Version=$(VERSION) \
	-X portal/internal/cli.VersionName=$(VERSION_NAME)"

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/portal

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean


run:
	go run ./cmd/portal $(ARGS)
