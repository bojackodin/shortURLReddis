MAIN_IMAGE := shortURLReddis:latest
MAIN_PACKAGE_PATH := ./cmd/web
BINARY_NAME := bin/web

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}
