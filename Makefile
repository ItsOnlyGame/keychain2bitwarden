# Variables
APP_NAME := keychain2bitwarden
BUILD_DIR := build
MAIN_FILE := ./cmd/keychain2bitwarden/main.go

ifeq ($(OS),Windows_NT)
	MKDIR := if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
	DEL := rmdir /s /q
	EXE_SUFFIX := .exe
else
	MKDIR := mkdir -p $(BUILD_DIR)
	DEL := rm -rf
	EXE_SUFFIX :=
endif

# Default target
all: build

# Build the application
build:
	@echo "Building $(APP_NAME) for GOOS=$(GOOS), GOARCH=$(GOARCH)..."
	@$(MKDIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME)$(EXE_SUFFIX) $(MAIN_FILE)

# Clean up build directory
clean:
	@echo "Cleaning up..."
	@$(DEL) $(BUILD_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
