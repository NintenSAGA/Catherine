CC := gcc

NAME := PushBlocks
SRC_DIRS := src
INC_DIRS := $(SRC_DIRS)/include
BLD_DIRS := build
OBJ_DIRS := $(patsubst %, $(BLD_DIRS)/%, $(SRC_DIRS))

SRCS := $(shell find $(SRC_DIRS) -name "*.c" -maxdepth 1)
DEPS := $(shell find $(INC_DIRS) -name "*.h" -maxdepth 1) $(SRCS)
OBJS := $(SRCS:.c=.o)

CFLAGS := -std=gnu11 -Wall -Werror

.PHONY: all run clean dir

.DEFAULT_GOAL := all

all: $(OBJS) $(DEPS)
	cd $(BLD_DIRS)\
	 && $(CC) $(OBJS) -o $(NAME)

$(OBJ_DIRS): 
	mkdir -p $@

%.o: %.c $(OBJ_DIRS)
	$(CC) $(CFLAGS) -c -o $(BLD_DIRS)/$@ $< 

run: all
	$(BLD_DIRS)/$(NAME)

clean:
	rm -rf $(BLD_DIRS)