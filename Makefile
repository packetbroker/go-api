# Copyright © 2019 The Things Industries B.V.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL = bash
GO = go
PROTOC = protoc

PBAPI ?= ../..
VERSION ?= v3
ROUTING_VERSION ?= v2
MAPPING_VERSION ?= v2
IAM_VERSION ?= v2
REPORTING_VERSION ?= v1

protos = $(wildcard $(PBAPI)/packetbroker/api/$(VERSION)/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/routing/$(ROUTING_VERSION)/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/mapping/$(MAPPING_VERSION)/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/iam/$(IAM_VERSION)/*.proto) \
	$(wildcard $(PBAPI)/packetbroker/api/reporting/$(REPORTING_VERSION)/*.proto)
prototargets = $(subst v1/,,$(patsubst $(PBAPI)/packetbroker/api/%.proto,%.pb.go,$(protos)))

openapis = $(PBAPI)/packetbroker/api/mapping/$(MAPPING_VERSION)/openapi.tmpl.json
openapitargets = $(subst v1/,,$(patsubst $(PBAPI)/packetbroker/api/%/openapi.tmpl.json,%/openapi/openapi.tmpl.json,$(openapis)))

.PHONY: all
all: $(prototargets)
all: $(openapitargets)

.PHONY: deps
deps:
	$(GO) get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	$(GO) get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

.PHONY: clean
clean:
	@rm $(prototargets)
	@rm $(openapitargets)

$(prototargets): $(protos)
	@set -e
	@mkdir -p build $(@D)
	@$(PROTOC) \
		--go_out="build" \
		--go-grpc_out="build" \
		--proto_path=$(PBAPI) $^
	@mv build/go.packetbroker.org/api/$(@D)/*.pb.go $(@D)/
	@rm -rf build

$(openapitargets): $(openapis)
	@cp $^ $@

# vim: ft=make
