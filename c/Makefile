CC     = gcc
CFLAGS = -std=c17 -Wall -Wextra -Wpedantic -g
INCL   = -Icommon

BIN    = build

.PHONY: all clean

all: $(BIN)/day01_part1

# Explicit rule for day01_part1
$(BIN)/day01_part1: day01/part1.c 
	mkdir -p $(BIN)
	$(CC) $(CFLAGS) $(INCL) $^ -o $@

clean:
	rm -rf $(BIN)
