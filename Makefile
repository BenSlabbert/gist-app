SHELL := /bin/bash

include .env
export

BIN="./bin"
SRC=$(shell find . -name "*.go")
GIT_COMMIT_ID = $(shell git log --format="%H" -n 1)
GIT_BRANCH_NAME = $(shell git symbolic-ref --short -q HEAD)

$(info BIN output dir: ${BIN})
$(info GIT_COMMIT_ID: ${GIT_COMMIT_ID})
$(info GIT_BRANCH_NAME: ${GIT_BRANCH_NAME})

.PHONY: web build upx fmt vet test mod install_deps clean

default: all

all: web clean install_deps mod vet fmt test build upx

web:
	$(MAKE) -C web

build:
	$(info ******************** build ********************)
	mkdir -p $(BIN)
	GOOS=linux ARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o bin/gist-app

# https://upx.github.io/
upx:
	$(info ******************** upx ********************)
	upx -o bin/gist-app.upx bin/gist-app

fmt:
	$(info ******************** fmt ********************)
	go fmt ./...
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

vet:
	$(info ******************** vet ********************)
	go vet ./...
	staticcheck ./...

test:
	$(info ******************** test ********************)
	go test ./...

mod:
	$(info ******************** mod ********************)
	go mod tidy
	go mod verify

install_deps:
	$(info ******************** install_deps ********************)
	go get -v ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest

clean:
	$(info ******************** clean ********************)
	rm -rf $(BIN)
