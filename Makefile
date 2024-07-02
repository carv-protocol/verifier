GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && cd cmd/verifier && go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/ ./...

.PHONY: build-all
# build for all environments
build-all:
	mkdir -p bin/windows/amd64 && env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/windows/amd64/ ./...
	mkdir -p bin/windows/386 && env GOOS=windows GOARCH=386 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/windows/386/ ./...
	mkdir -p bin/macos/amd64 && env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/macos/amd64/ ./...
	mkdir -p bin/linux/386 && env GOOS=linux GOARCH=386 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/linux/386/ ./...
	mkdir -p bin/linux/amd64 && env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/linux/amd64/ ./...
	mkdir -p bin/linux/arm && env GOOS=linux GOARCH=arm go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/linux/arm/ ./...
	mkdir -p bin/freebsd/386 && env GOOS=freebsd GOARCH=386 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/freebsd/386/ ./...
	mkdir -p bin/freebsd/amd64 && env GOOS=freebsd GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/freebsd/amd64/ ./...
	mkdir -p bin/freebsd/arm && env GOOS=freebsd GOARCH=arm go build -ldflags "-X main.Version=$(VERSION)" -o ../../bin/freebsd/arm/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;
	make build;

# wasm build
.PHONY: wasm
wasm:
	cd cmd/wasm && GOOS=js GOARCH=wasm go build -o main.wasm
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
