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
DOCKER = docker

PROTO_PATH ?= ../..
VERSION ?= v3
ROUTING_VERSION ?= v1
MAPPING_VERSION ?= v1
IAM_VERSION ?= v1
PROTOC_IMAGE = thethingsindustries/protoc:3.1.28-tts

PROTOC = $(DOCKER) run --rm \
	--mount type=bind,src=$(abspath $(PROTO_PATH))/packetbroker,dst=/src/packetbroker \
	--mount type=bind,src=$(PWD),dst=/out/go.packetbroker.org/api \
	$(PROTOC_IMAGE) \
	-I/src \
	--gogottn_out=plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:/out

protos = $(wildcard $(PROTO_PATH)/packetbroker/api/$(VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/routing/$(ROUTING_VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/mapping/$(MAPPING_VERSION)/*.proto) \
	$(wildcard $(PROTO_PATH)/packetbroker/api/iam/$(IAM_VERSION)/*.proto)
targets = $(subst v1/,,$(patsubst $(PROTO_PATH)/packetbroker/api/%.proto,%.pb.go,$(protos)))

.PHONY: all
all: $(targets)

.PHONY: deps
deps:
	$(DOCKER) pull $(PROTOC_IMAGE)

.PHONY: clean
clean:
	@rm $(targets)

$(targets): $(protos)
	$(PROTOC) /src/packetbroker/api/$(VERSION)/*.proto
	$(PROTOC) /src/packetbroker/api/routing/$(ROUTING_VERSION)/*.proto
	$(PROTOC) /src/packetbroker/api/mapping/$(MAPPING_VERSION)/*.proto
	$(PROTOC) /src/packetbroker/api/iam/$(IAM_VERSION)/*.proto

# vim: ft=make
