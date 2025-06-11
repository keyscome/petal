# Makefile for Petal multi-platform builds

APP_NAME := petal
BUILD_DIR := build
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

.PHONY: all build clean dist

all: build

build:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		IFS="/" read -r GOOS GOARCH <<< "$$platform"; \
		output="$(APP_NAME)-$$GOOS-$$GOARCH"; \
		if [ "$$GOOS" = "windows" ]; then \
			output="$$output.exe"; \
		fi; \
		echo "🔧 Building $$platform -> $$output"; \
		CGO_ENABLED=0 GOOS=$$GOOS GOARCH=$$GOARCH go build -o "$(BUILD_DIR)/$$output" .; \
	done

clean:
	@rm -rf $(BUILD_DIR)
	@echo "🧹 Cleaned $(BUILD_DIR)"

dist: build
	@echo "✅ All builds are in $(BUILD_DIR)"

# Usage:
#   make        # build all
#   make clean  # clean build artifacts
#   make dist   # build + summary
