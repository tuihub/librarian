PACKAGE_NAME := github.com/tuihub/librarian
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PROTO_VERSION=$(shell go list -m -f '{{.Version}}' github.com/tuihub/protos)
GOLANG_CROSS_VERSION ?= v1.20.7
SHELL:=/bin/bash

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1

init-test:
	cd tests && make init

.PHONY: generate
# generate code
generate: generate-config generate-code

generate-config:
	protoc --proto_path=./internal \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

generate-code:
	go get github.com/google/wire/cmd/wire@latest
	go get github.com/jmattheis/goverter@v1.2.0
	go generate ./...
	go mod tidy

.PHONY: lint
# lint files
lint:
	golangci-lint run --fix

.PHONY: test-unit
# run go test
test-unit:
	go test -coverpkg=./... -race -coverprofile=coverage-unit.out -covermode=atomic ./...

.PHONY: test-goc
# run goc test
test-goc:
	cd tests && make all

test-postprocess:
	@while read -r p || [ -n "$$p" ]; \
	do \
	if [[ "$(GOHOSTOS)" == "darwin" ]]; then \
	  sed -i '' "/$${p//\//\\/}/d" ./coverage-unit.out; \
	  sed -i '' "/$${p//\//\\/}/d" ./coverage-goc.out; \
	else \
	  sed -i "/$${p//\//\\/}/d" ./coverage-unit.out; \
	  sed -i "/$${p//\//\\/}/d" ./coverage-goc.out; \
	fi \
	done < ./.coverageignore

test-all: test-unit test-goc test-postprocess

.PHONY: run
# run server
run:
	CGO_ENABLED=1 kratos run

.PHONY: build
# build server in debug mode
build:
	mkdir -p bin/ && go build -tags debug -ldflags "-X main.version=$(VERSION) -X main.protoVersion=$(PROTO_VERSION)" -o ./bin/ ./...

.PHONY: release-dry-run
# build server in release mode, for manual test
release-dry-run:
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-e PROTO_VERSION=$(PROTO_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --clean --skip-validate --skip-publish

.PHONY: release
# build server in release mode, for CI, do not run manually
release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-e PROTO_VERSION=$(PROTO_VERSION) \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release

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
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
