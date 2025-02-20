SRC_DIR := cmd
BIN_DIR := bin

CMDS := jokenpo

CGO_ENABLED := 0
C_FLAGS := -s -w

.PHONY: all stage build clean vet fmt test race run

all: clean stage vet fmt test race build run

stage:
	@echo "Creating build folder"
	mkdir -p $(BIN_DIR)


build:
	@echo "Compiling $(CMDS).... "
	@go build -ldflags "$(C_FLAGS)" -o $(BIN_DIR)/$(CMDS) $(SRC_DIR)/main.go

.PHONY:

test:
	@echo "Testing"
	@go test ./...

race:
	@echo "Testing race condition"
	@go test -race ./...

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
