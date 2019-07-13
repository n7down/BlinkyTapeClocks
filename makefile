VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MAKEFLAGS += --silent
PID := /tmp/.$(PROJECTNAME).pid
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(GOPATH)/src/github.com/n7down/timelord/cmd/timelord/*.go
ALLFILES=$(shell find . -name '*.go')

install:
	echo "installing... \c"
	@go get ./... 
	echo "done"

build: clean
	echo "building... \c"
	@GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	echo "done"

generate:
	echo "generating dependency files... \c"
	@GOBIN=$(GOBIN) go generate ./...
	echo "done"

compile: install build

start-server: stop-server
	echo "starting server... \c"
	@$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	echo "done"
	cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	echo "stopping server... \c"
	@touch $(PID)
	@kill `cat $(PID)` 2> /dev/null || true
	@rm $(PID)
	echo "done"

start: compile start-server

stop: stop-server

test:
		@go test -short ${ALLFILES}

vet:
		@go vet ${ALLFILES}

lint:
		@for file in ${ALLFILES); do \
			golint $$file ; \
		done

clean:
	echo "cleaning build cache... \c"
	@go clean
	@rm -rf bin/
	echo "done"

update:
	@clear
	make stop
	git pull origin dev
	make start

help:
	echo "Choose a command run in $(PROJECTNAME):"
	echo " install - installs all dependencies for the project"
	echo " build - builds a binary"
	echo " compile - installs all dependencies and builds a binary"
	echo " clean - cleans the cache and cleans up the build files"
