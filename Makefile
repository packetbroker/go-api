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
GOPATH ?= $(shell $(GO) env GOPATH)

targets = $(addsuffix /packetbroker.pb.go, $(VERSION))

.PHONY: all
all: $(targets)

.PHONY: deps
deps:
	$(GO) get github.com/gogo/protobuf/protoc-gen-gogofaster
	$(GO) get google.golang.org/grpc

.PHONY: clean
clean:
	@rm $(VERSION)/*.pb.go

$(VERSION):
	@mkdir -p $@

$(targets): $(VERSION)/%.pb.go: $(VERSION)
	$(PROTOC) -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogofaster_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:../../ \
		--proto_path=$(PROTO_PATH) \
		$(PROTO_PATH)/packetbroker/api/$(@D)/*.proto

# vim: ft=make
