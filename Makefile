.DEFAULT_GOAL := all
.PHONY: all install clean dist dist-prepare dist-osx dist-linux

REPO := github.com/unbounce/stack-toolkit
DIST_DIR := $(GOPATH)/dist

all: install

install:
	go install $(REPO)/cli/stacks
	go install $(REPO)/cli/stack-instances

clean:
	rm -rf $(DIST_DIR)

dist-prepare:
ifndef VERSION
	@echo "VERSION is required to continue" && exit 1
endif
	mkdir -p $(DIST_DIR)

dist: dist-prepare dist-osx dist-linux

dist-osx:
	$(eval DARWIN_AMD64_DIST_DIR := $(DIST_DIR)/darwin.amd64)
	mkdir -p $(DARWIN_AMD64_DIST_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(DARWIN_AMD64_DIST_DIR)/stacks $(REPO)/cli/stacks
	GOOS=darwin GOARCH=amd64 go build -o $(DARWIN_AMD64_DIST_DIR)/stack-instances $(REPO)/cli/stack-instances
	tar cfz $(DIST_DIR)/stack-toolkit.$(VERSION).darwin-amd64.tgz -C $(DARWIN_AMD64_DIST_DIR) .

dist-linux:
	$(eval LINUX_AMD64_DIST_DIR := $(DIST_DIR)/linux.amd64)
	mkdir -p $(LINUX_AMD64_DIST_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_AMD64_DIST_DIR)/stacks $(REPO)/cli/stacks
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_AMD64_DIST_DIR)/stack-instances $(REPO)/cli/stack-instances
	tar cfz $(DIST_DIR)/stack-toolkit.$(VERSION).linux-amd64.tgz -C $(LINUX_AMD64_DIST_DIR) .

