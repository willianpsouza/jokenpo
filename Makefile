SRC_DIR := cmd
BIN_DIR := bin

CMDS := jokenpo

CGO_ENABLED := 0
C_FLAGS := -s -w

.PHONY: all stage build clean vet fmt test run

all: clean stage vet fmt test build run

stage:
	@echo "Creating build folder"
	mkdir -p $(BIN_DIR)


build:
	@echo "Compiling $(CMDS).... "
	@go build -ldflags "$(C_FLAGS)" -o $(BIN_DIR)/$(CMDS) main.go

.PHONY:

test:
	@echo "Testing"
	@go test ./...

fmt:
	@echo "Formating"
	@go fmt ./...

vet:
	@echo "Vetting"
	@go vet ./...

clean:
	@echo "Clean all"
	@rm -rf $(BIN_DIR)

run:
	@echo "Running game"
	@$(BIN_DIR)/$(CMDS)