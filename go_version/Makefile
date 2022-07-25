
NAME		:= Catherine

GO			:= go
MAIN		:= src/main.go
BIN_DIR 	:= bin

.PHONY: all run clean dir

.DEFAULT_GOAL := all

all: $(MAIN)
	$(GO) build -o $(BIN_DIR)/$(NAME)  $(MAIN)

build-docker: $(MAIN)
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.18\
	 bash -c '\
	 	for GOOS in darwin linux; do\
	 		for GOARCH in amd64; do\
	 			export GOOS GOARCH; \
				$(GO) build -o $(BIN_DIR)/$(NAME)-$$GOOS-$$GOARCH $(MAIN); \
			done \
		done \
	 '

run: all
	make run_only

run_only:
	$(BIN_DIR)/$(NAME)

clean: $(BIN_DIR)
	rm $(BIN_DIR)/*

