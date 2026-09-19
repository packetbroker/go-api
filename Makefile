# SPDX-FileCopyrightText: Copyright 2019 The Things Industries B.V.
# SPDX-License-Identifier: Apache-2.0

SHELL = bash
# The repository may sit in a Go workspace of the caller; build and run tools against this module only.
export GOWORK = off
GO = go
GIT = git
CURL = curl

# Tools (buf, protoc plugins, linters) are tool dependencies of the tools module, not of the public module.
GOTOOL = $(GO) tool -modfile=tools/go.mod

# Commit of github.com/packetbroker/api that the checked-in code is generated from.
PBAPI_REF = 1ece72d9bc7605cd1d942e75d70baf7007056841
# buf input to generate from. Defaults to the API repository at PBAPI_REF, which buf fetches itself.
# Override with a local checkout of the API repository, e.g. PBAPI_INPUT=../api
PBAPI_INPUT ?= https://github.com/packetbroker/api.git\#ref=$(PBAPI_REF)

# Packages with generated code: <api dir in packetbroker/api>:<package dir in this repository>.
APIS = v3:v3 routing/v1:routing routing/v2:routing/v2 iam/v1:iam iam/v2:iam/v2 mapping/v2:mapping/v2 reporting/v1:reporting

# OpenAPI templates are not protos: <api dir in packetbroker/api>:<target in this repository>.
openapis = mapping/v2:mapping/v2/openapi/openapi.tmpl.json

.PHONY: all
all: generate openapi

.PHONY: clean
clean:
	@for a in $(APIS); do rm -f $${a#*:}/*.pb.go; done
	@for o in $(openapis); do rm -f $${o#*:}; done

# buf.gen.yaml maps the APIs to Go packages (managed mode) and writes them to their directories.
.PHONY: generate
generate:
	@$(GOTOOL) buf generate $(PBAPI_INPUT)

.PHONY: openapi
openapi:
	@for o in $(openapis); do \
		src=packetbroker/api/$${o%%:*}/openapi.tmpl.json; dst=$${o#*:}; \
		mkdir -p $$(dirname $$dst); \
		if [ -d "$(PBAPI_INPUT)" ]; then \
			cp "$(PBAPI_INPUT)/$$src" $$dst || exit 1; \
		else \
			$(CURL) -fsSL "https://raw.githubusercontent.com/packetbroker/api/$(PBAPI_REF)/$$src" -o $$dst || exit 1; \
		fi; \
	done

.PHONY: build
build:
	@$(GO) build ./...

.PHONY: test
test:
	@$(GO) test -race -covermode=atomic ./...

.PHONY: deps.tidy
deps.tidy:
	@$(GO) mod tidy
	@cd tools && $(GO) mod tidy

.PHONY: fmt
fmt:
	@$(GOTOOL) gofumpt -w -extra -l .

.PHONY: quality
quality:
	@$(GOTOOL) golangci-lint run --timeout 5m0s --allow-parallel-runners --max-issues-per-linter 0 --max-same-issues 0 $(GO_LINT_FLAGS) ./...

BASE_REF ?= master
.PHONY: quality.new
quality.new:
	@$(MAKE) quality GO_LINT_FLAGS="$(strip $(GO_LINT_FLAGS) --new-from-rev=origin/$(BASE_REF))"

.PHONY: git.nodiff
git.nodiff:
	@$(GIT) diff --exit-code

# vim: ft=make
