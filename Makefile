CC = gcc
CFLAGS = -Wall -Wextra -pedantic
SRCS := $(wildcard  leetcode/*/*.c) 
OBJECTS := $(SRCS:.c=.o)

all: $(OBJECTS)

leetcode/%/solution.o: leetcode/%/solution.c
	$(CC) $(CFLAGS) -o $@ $<

clean:
	@rm -f $(OBJECTS) 

.PHONY: clean
