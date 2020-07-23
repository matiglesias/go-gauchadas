EXEC=./main

SHELL=/bin/bash

all: run
SRCS=$(shell find . -type f -name '*.go')

build:
$(EXEC): $(SRCS)
	go build ./api/main.go

run: $(EXEC)
	source .env && $(EXEC)

clean:
	rm -f $(EXEC)

test:
	go test -v ./...
