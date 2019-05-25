NAME := terraform-provider-yaml
PLATFORMS ?= darwin/amd64 linux/amd64 windows/amd64
VERSION = $(shell git describe 1>/dev/null 2>/dev/null && echo "_$$(git describe)")

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

BASE := $(NAME)$(VERSION)
RELEASE_DIR := ./release

all: clean test release

clean:
	rm -rf $(RELEASE_DIR) ./$(NAME)*

format:
	GOPROXY="off" GOFLAGS="-mod=vendor" go fmt ./...

test:
	GOPROXY="off" GOFLAGS="-mod=vendor" go test -v ./...
	GOPROXY="off" GOFLAGS="-mod=vendor" go vet ./...

build:
	GOPROXY="off" GOFLAGS="-mod=vendor" go build -o $(BASE)

release: $(PLATFORMS)

$(PLATFORMS):
	GOPROXY="off" GOFLAGS="-mod=vendor" GOOS=$(os) GOARCH=$(arch) go build -o '$(RELEASE_DIR)/$(BASE)-$(os)-$(arch)'

.PHONY: $(PLATFORMS) release build test fmt clean all
