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

PROTO_PATH ?= ../..
VERSION ?= v3
ROUTING_VERSION ?= v1
MAPPING_VERSION ?= v1
IAM_VERSION ?= v1

protos = $(wildcard $(PROTO_PATH)/packetbroker/api/$(VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/routing/$(ROUTING_VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/mapping/$(MAPPING_VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/iam/$(IAM_VERSION)/*.proto)
targets = $(subst v1/,,$(patsubst $(PROTO_PATH)/packetbroker/api/%.proto,%.pb.go,$(protos)))

.PHONY: all
all: $(targets)

.PHONY: deps
deps:
	$(GO) get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
	$(GO) get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

.PHONY: clean
clean:
	@rm $(targets)

$(targets): $(protos)
	@mkdir -p build $(@D)
	@$(PROTOC) \
		--go_out="build" \
		--go-grpc_out="build" \
		--proto_path=$(PROTO_PATH) $^
	@mv build/go.packetbroker.org/api/$(@D)/*.pb.go $(@D)/

# vim: ft=make
