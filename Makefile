.DEFAULT_GOAL := all
.PHONY: all install clean dist dist-prepare dist-osx dist-linux deps help

PROJECT_NAME := stack-toolkit

REPO := github.com/unbounce/stack-toolkit
DIST_DIR := $(GOPATH)/dist

LDFLAGS=-ldflags="-X main.Version=$(VERSION)"
ARCH_DIST_DIR=$(DIST_DIR)/$(OS)-$(ARCH)

BUILD_COMMAND = GOOS=$(OS) GOARCH=$(ARCH) go build $(LDFLAGS) -o $(ARCH_DIST_DIR)/$(BINARY) $(REPO)/cli/$(BINARY)
PKG_COMMAND = tar cfz $(DIST_DIR)/$(PROJECT_NAME).$(VERSION).$(OS)-$(ARCH).tar.gz -C $(ARCH_DIST_DIR) .
MKDIR_COMMAND = mkdir -p $(DIST_DIR)/$(OS)-$(ARCH)

all: install

help: ## displays this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

deps: ## Installs third-party dependencies
	go get -u github.com/urfave/cli
	go get -u github.com/aws/aws-sdk-go

install: ## Builds and installs the binaries locally
	go install $(REPO)/cli/stacks
	go install $(REPO)/cli/stack-instances

clean: ## Removes any derived files and directories
	rm -rf $(DIST_DIR)

dist-prepare:
ifndef VERSION
	$(error VERSION is required to continue)
endif
	mkdir -p $(DIST_DIR)

dist: dist-prepare dist-osx dist-linux ## Creates artifacts for all OSes

dist-osx:
	$(eval OS := darwin)
	$(eval ARCH := amd64)
	$(MKDIR_COMMAND)

	$(eval BINARY := stacks)
	$(BUILD_COMMAND)

	$(eval BINARY := stack-instances)
	$(BUILD_COMMAND)

	$(PKG_COMMAND)

dist-linux:
	$(eval OS := linux)
	$(eval ARCH := amd64)
	$(MKDIR_COMMAND)

	$(eval BINARY := stacks)
	$(BUILD_COMMAND)

	$(eval BINARY := stack-instances)
	$(BUILD_COMMAND)

	$(PKG_COMMAND)

