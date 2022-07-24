
NAME		:= PushBlock

GO			:= go
MAIN		:= src/main.go
BIN_DIR 	:= bin

.PHONY: all run clean dir

.DEFAULT_GOAL := all

all: $(MAIN)
	$(GO) build -o $(BIN_DIR)/$(NAME) $(MAIN)

run: all
	make run_only

run_only:
	$(BIN_DIR)/$(NAME)

clean: $(BIN_DIR)
	rm $(BIN_DIR)/*

