# SPDX-FileCopyrightText: Copyright 2019 The Things Industries B.V.
# SPDX-License-Identifier: Apache-2.0

SHELL = bash
GO = go
GIT = git
PROTOC = protoc

# Version of protoc that the checked-in code is generated with. CI installs exactly this version.
PROTOC_VERSION = 36.1
# Commit of github.com/packetbroker/api that the checked-in code is generated from.
PBAPI_REF = 50a372db610df3e503457e9ddbc6334c50519412
# Directory that contains the packetbroker/api checkout, i.e. protos live in $(PBAPI)/packetbroker/api.
PBAPI ?= ../..

# Go modules in this repository, excluding the root module that only carries tools.
MODULES = v3 routing routing/v2 iam iam/v2 mapping/v2 reporting

# protoc plugins come from the tool dependencies of the root module.
PROTOC_GEN_GO = $(shell $(GO) tool -n protoc-gen-go)
PROTOC_GEN_GO_GRPC = $(shell $(GO) tool -n protoc-gen-go-grpc)

protos = $(wildcard $(PBAPI)/packetbroker/api/v3/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/routing/v1/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/routing/v2/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/mapping/v2/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/iam/v1/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/iam/v2/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/reporting/v1/*.proto)
prototargets = $(subst v1/,,$(patsubst $(PBAPI)/packetbroker/api/%.proto,%.pb.go,$(protos)))

openapis = $(PBAPI)/packetbroker/api/mapping/v2/openapi.tmpl.json
openapitargets = $(subst v1/,,$(patsubst $(PBAPI)/packetbroker/api/%/openapi.tmpl.json,%/openapi/openapi.tmpl.json,$(openapis)))

.PHONY: all
all: $(prototargets)
all: $(openapitargets)

.PHONY: clean
clean:
	@rm -f $(prototargets)
	@rm -f $(openapitargets)

$(prototargets): $(protos)
	@set -e
	@mkdir -p build $(@D)
	@$(PROTOC) \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go_out="build" \
		--go-grpc_out="build" \
		--proto_path=$(PBAPI) $^
	@mv build/go.packetbroker.org/api/$(@D)/*.pb.go $(@D)/
	@rm -rf build

$(openapitargets): $(openapis)
	@cp $^ $@

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
		(cd $$m && $(GO) tool golangci-lint run --timeout 5m0s --allow-parallel-runners $(GO_LINT_FLAGS) ./...) || exit 1; \
	done

BASE_REF ?= master
.PHONY: quality.new
quality.new:
	@$(MAKE) quality GO_LINT_FLAGS="$(strip $(GO_LINT_FLAGS) --new-from-rev=origin/$(BASE_REF))"

.PHONY: git.nodiff
git.nodiff:
	@$(GIT) diff --exit-code

# vim: ft=make
