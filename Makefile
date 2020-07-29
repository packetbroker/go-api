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

targets = $(addsuffix /packetbroker.pb.go, $(VERSION))

.PHONY: all
all: $(targets)

.PHONY: deps
deps:
	$(GO) get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
	$(GO) get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@1f47ba4663831f2a9c28a62a7de3ff8bc45078f0

.PHONY: clean
clean:
	@rm $(VERSION)/*.pb.go

$(VERSION):
	@mkdir -p $@

$(targets): $(VERSION)/%.pb.go: $(VERSION)
	mkdir -p build
	$(PROTOC) \
		--go_out="build" \
		--go-grpc_out="build" \
		--proto_path=$(PROTO_PATH) \
		$(PROTO_PATH)/packetbroker/api/$(@D)/*.proto
	mv build/go.packetbroker.org/api/$(VERSION)/*.pb.go $(VERSION)/

# vim: ft=make
