# SPDX-FileCopyrightText: Copyright 2019 The Things Industries B.V.
# SPDX-License-Identifier: Apache-2.0

SHELL = bash
GO = go
GIT = git
CURL = curl

# Commit of github.com/packetbroker/api that the checked-in code is generated from.
PBAPI_REF = 1ece72d9bc7605cd1d942e75d70baf7007056841
# buf input to generate from. Defaults to the API repository at PBAPI_REF, which buf fetches itself.
# Override with a local checkout of the API repository, e.g. PBAPI_INPUT=../api
PBAPI_INPUT ?= https://github.com/packetbroker/api.git\#ref=$(PBAPI_REF)

# Go modules in this repository, excluding the root module that only carries tools.
# buf writes the generated code per go_package to build/go.packetbroker.org/api/<module>.
MODULES = v3 routing routing/v2 iam iam/v2 mapping/v2 reporting

# OpenAPI templates are not protos: <api dir in packetbroker/api>:<target in this repository>.
openapis = mapping/v2:mapping/v2/openapi/openapi.tmpl.json

.PHONY: all
all: generate openapi

.PHONY: clean
clean:
	@for m in $(MODULES); do rm -f $$m/*.pb.go; done
	@for o in $(openapis); do rm -f $${o#*:}; done
	@rm -rf build

# buf and the plugins are tools of the root module; GOWORK=off keeps the workspace module graph out of the tool build.
.PHONY: generate
generate:
	@rm -rf build
	@GOWORK=off $(GO) tool buf generate $(PBAPI_INPUT)
	@for m in $(MODULES); do mv build/go.packetbroker.org/api/$$m/*.pb.go $$m/ || exit 1; done
	@rm -rf build

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
	@for m in $(MODULES); do \
		echo "build $$m"; \
		(cd $$m && GOWORK=off $(GO) build ./...) || exit 1; \
	done

.PHONY: test
test:
	@for m in $(MODULES); do \
		echo "test $$m"; \
		(cd $$m && GOWORK=off $(GO) test -race -covermode=atomic ./...) || exit 1; \
	done

.PHONY: deps.tidy
deps.tidy:
	@for m in . $(MODULES); do \
		echo "tidy $$m"; \
		(cd $$m && $(GO) mod tidy) || exit 1; \
	done

.PHONY: fmt
fmt:
	@$(GO) tool gofumpt -w -extra -l .

.PHONY: quality
quality:
	@for m in $(MODULES); do \
		echo "lint $$m"; \
		(cd $$m && $(GO) tool golangci-lint run --timeout 5m0s --allow-parallel-runners --max-issues-per-linter 0 --max-same-issues 0 $(GO_LINT_FLAGS) ./...) || exit 1; \
	done

BASE_REF ?= master
.PHONY: quality.new
quality.new:
	@$(MAKE) quality GO_LINT_FLAGS="$(strip $(GO_LINT_FLAGS) --new-from-rev=origin/$(BASE_REF))"

.PHONY: git.nodiff
git.nodiff:
	@$(GIT) diff --exit-code

# vim: ft=make
