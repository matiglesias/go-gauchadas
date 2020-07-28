EXEC=./main

SHELL=/bin/bash

all: run
SRCS=$(shell find . -type f -name '*.go')

db:
	docker run -d -v mongo-data:/data/db -v mongo-config:/data/configdb -p:27017:27017 --name test-db mongo:latest

db-stop:
	docker stop test-db && docker rm test-db

build:
$(EXEC): $(SRCS)
	go build ./api/main.go

run: $(EXEC)
	source .env && $(EXEC)

clean:
	rm -f $(EXEC)

test:
	go test -v ./...
