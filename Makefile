NAME := terraform-provider-yaml
PLATFORMS ?= darwin/amd64 linux/amd64 windows/amd64
VERSION ?= $(shell git describe &>/dev/null && echo "_$$(git describe)")

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

BASE := $(NAME)$(VERSION)
RELEASE_DIR := ./release

all: clean format test release

clean:
	rm -rf $(RELEASE_DIR) ./$(BASE)*

format:
	go fmt ./...

test:
	go test -v ./...
	go vet ./...

build:
	go build -o $(BASE)

release: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o '$(RELEASE_DIR)/$(BASE)-$(os)-$(arch)'

.PHONY: $(PLATFORMS) release build test fmt clean all
