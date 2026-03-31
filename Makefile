# Makefile for Petal multi-platform builds

APP_NAME := petal
BUILD_DIR := build
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64

.PHONY: all build clean dist examples docker-up docker-down

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
	@obsutil cp "$(BUILD_DIR)/petal-linux-amd64" obs://selfhosted/petal/
	@obsutil cp "$(BUILD_DIR)/petal-windows-amd64.exe" obs://selfhosted/petal/
	@obsutil cp "$(BUILD_DIR)/petal-darwin-amd64" obs://selfhosted/petal/
	@obsutil cp "$(BUILD_DIR)/petal-darwin-arm64" obs://selfhosted/petal/
	@echo "✅ All builds uploaded to OBS"
	
dist-examples:
	@echo "📦 Building example files"
	@tar zcf examples.tgz examples
	@echo "📦 Uploading examples to OBS"
	@obsutil cp examples.tgz obs://selfhosted/petal/
	@rm -f examples.tgz
	@echo "✅ Example files uploaded to OBS"

# Spin up three Docker-based SSH test nodes (node1, node2, node3) on localhost
docker-up:
	@bash docker/setup.sh

# Tear down the Docker test nodes
docker-down:
	@bash docker/setup.sh --down

# Usage:
#   make              # build all
#   make clean        # clean build artifacts
#   make dist         # build + upload to OBS
#   make docker-up    # start local SSH test cluster via Docker
#   make docker-down  # stop local SSH test cluster

