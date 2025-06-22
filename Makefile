BUILD_OPTS:=-ldflags="-s -w"

# Absolute path to the root directory of the project
ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Absolute path to this makefile
THIS_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))

# Configure dev utilities to install into a local folder to avoid clobbering global installs
TOOLS_FILE := tools.go
TOOLS_ABS_PATH := $(abspath tools)

BINARY_NAME_STRAVA_APP := strava-app

export GOBIN := $(TOOLS_ABS_PATH)
export PATH := $(TOOLS_ABS_PATH):$(PATH)
export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)
export GO111MODULE=on

clean:
	go clean -i
	rm -rf vendor/
	rm -rf tools/
	go mod tidy
	go mod vendor

# Ensure a golang-installed cli tool is installed
# Usage: tools/<toolname>
# Parses $(TOOLS_FILE) for a package import statement ending in the
# given toolname and if found runs `go install` on that package
tools/%: SHELL:=/bin/bash
tools/%:
	$(eval toolpkg = $(shell cat tools.go | grep -Ei  '^\s*_ ".+/$*.*"$$' | sed 's/.*_ "\(.*\)".*/\1/'))
	@if [[ -z "$(toolpkg)" ]]; then \
		echo ""; \
		echo "ERROR: Tool '$*' is not known."; \
		echo "       If it is a go tool, add it to $(TOOLS_FILE)"; \
		echo ""; \
		exit 1; \
	fi
	$(eval toolpkg := $(toolpkg)@$(shell go list -f '{{.Module.Version}}' -find $(toolpkg)))
	go install $(toolpkg)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_OPTS) -o $(BINARY_NAME_STRAVA_APP) ./cmd/

lint: tools/golangci-lint
	golangci-lint run -c .golangci.yml --timeout 5m

## Generate sources from OpenAPI spec
gen.openapi: tools/oapi-codegen
	oapi-codegen -config internal/strava/web/openapi/server.gen.yml internal/strava/web/openapi/spec.yml

.PRECIOUS: .env tools/%
