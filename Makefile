VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
PROJECT_NAME := goload

.PHONY: build 
build:
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME) \
		cmd/$(PROJECT_NAME)/*.go