BINARY=cgotest

CFLAGS=-pipe
LDFLAGS=-ldl

SOURCES := $(shell ls *.c)

.DEFAULT_GOAL: all

lib:
	go build -v -buildmode=c-shared -o=dynlib.so ./dynlib

$(BINARY):
	gcc $(CFLAGS) -o $(BINARY) $(SOURCES) $(LDFLAGS)

all: lib $(BINARY)
