// Copyright © 2021 The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: packetbroker/api/routing/v2/service.proto

package routingpb

import (
	v3 "go.packetbroker.org/api/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListUplinkRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of items to skip.
	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of items.
	// If 0, use the server's default.
	// The actual limit may be capped by the server.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListUplinkRoutesRequest) Reset() {
	*x = ListUplinkRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUplinkRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUplinkRoutesRequest) ProtoMessage() {}

func (x *ListUplinkRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUplinkRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListUplinkRoutesRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListUplinkRoutesRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUplinkRoutesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListUplinkRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*v3.DevAddrPrefixRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Total number of routes.
	Total uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListUplinkRoutesResponse) Reset() {
	*x = ListUplinkRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUplinkRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUplinkRoutesResponse) ProtoMessage() {}

func (x *ListUplinkRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUplinkRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListUplinkRoutesResponse) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListUplinkRoutesResponse) GetRoutes() []*v3.DevAddrPrefixRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ListUplinkRoutesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListJoinRequestRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of items to skip.
	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of items.
	// If 0, use the server's default.
	// The actual limit may be capped by the server.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListJoinRequestRoutesRequest) Reset() {
	*x = ListJoinRequestRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJoinRequestRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJoinRequestRoutesRequest) ProtoMessage() {}

func (x *ListJoinRequestRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJoinRequestRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListJoinRequestRoutesRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListJoinRequestRoutesRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListJoinRequestRoutesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListJoinRequestRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*v3.JoinEUIPrefixRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Total number of routes.
	Total uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListJoinRequestRoutesResponse) Reset() {
	*x = ListJoinRequestRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJoinRequestRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJoinRequestRoutesResponse) ProtoMessage() {}

func (x *ListJoinRequestRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJoinRequestRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListJoinRequestRoutesResponse) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListJoinRequestRoutesResponse) GetRoutes() []*v3.JoinEUIPrefixRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ListJoinRequestRoutesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListNetworkTargetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of items to skip.
	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of items.
	// If 0, use the server's default.
	// The actual limit may be capped by the server.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListNetworkTargetsRequest) Reset() {
	*x = ListNetworkTargetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworkTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkTargetsRequest) ProtoMessage() {}

func (x *ListNetworkTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkTargetsRequest.ProtoReflect.Descriptor instead.
func (*ListNetworkTargetsRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListNetworkTargetsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListNetworkTargetsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListNetworkTargetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*v3.NetworkTarget `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	// Total number of targets.
	Total uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListNetworkTargetsResponse) Reset() {
	*x = ListNetworkTargetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworkTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkTargetsResponse) ProtoMessage() {}

func (x *ListNetworkTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkTargetsResponse.ProtoReflect.Descriptor instead.
func (*ListNetworkTargetsResponse) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListNetworkTargetsResponse) GetTargets() []*v3.NetworkTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *ListNetworkTargetsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type TopicSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Topic name.
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// Subscription group.
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *TopicSubscription) Reset() {
	*x = TopicSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicSubscription) ProtoMessage() {}

func (x *TopicSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicSubscription.ProtoReflect.Descriptor instead.
func (*TopicSubscription) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{6}
}

func (x *TopicSubscription) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *TopicSubscription) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type TopicSubscriptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*TopicSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *TopicSubscriptions) Reset() {
	*x = TopicSubscriptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicSubscriptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicSubscriptions) ProtoMessage() {}

func (x *TopicSubscriptions) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicSubscriptions.ProtoReflect.Descriptor instead.
func (*TopicSubscriptions) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{7}
}

func (x *TopicSubscriptions) GetSubscriptions() []*TopicSubscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type TopicSubscriptionsChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Change:
	//
	//	*TopicSubscriptionsChange_Subscribes
	//	*TopicSubscriptionsChange_Unsubscribes
	Change isTopicSubscriptionsChange_Change `protobuf_oneof:"change"`
}

func (x *TopicSubscriptionsChange) Reset() {
	*x = TopicSubscriptionsChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicSubscriptionsChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicSubscriptionsChange) ProtoMessage() {}

func (x *TopicSubscriptionsChange) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicSubscriptionsChange.ProtoReflect.Descriptor instead.
func (*TopicSubscriptionsChange) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{8}
}

func (m *TopicSubscriptionsChange) GetChange() isTopicSubscriptionsChange_Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func (x *TopicSubscriptionsChange) GetSubscribes() *TopicSubscriptions {
	if x, ok := x.GetChange().(*TopicSubscriptionsChange_Subscribes); ok {
		return x.Subscribes
	}
	return nil
}

func (x *TopicSubscriptionsChange) GetUnsubscribes() *TopicSubscriptions {
	if x, ok := x.GetChange().(*TopicSubscriptionsChange_Unsubscribes); ok {
		return x.Unsubscribes
	}
	return nil
}

type isTopicSubscriptionsChange_Change interface {
	isTopicSubscriptionsChange_Change()
}

type TopicSubscriptionsChange_Subscribes struct {
	Subscribes *TopicSubscriptions `protobuf:"bytes,1,opt,name=subscribes,proto3,oneof"`
}

type TopicSubscriptionsChange_Unsubscribes struct {
	Unsubscribes *TopicSubscriptions `protobuf:"bytes,2,opt,name=unsubscribes,proto3,oneof"`
}

func (*TopicSubscriptionsChange_Subscribes) isTopicSubscriptionsChange_Change() {}

func (*TopicSubscriptionsChange_Unsubscribes) isTopicSubscriptionsChange_Change() {}

type TopicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Topic name.
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// Message body.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Message metadata.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TopicMessage) Reset() {
	*x = TopicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMessage) ProtoMessage() {}

func (x *TopicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_routing_v2_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMessage.ProtoReflect.Descriptor instead.
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_routing_v2_service_proto_rawDescGZIP(), []int{9}
}

func (x *TopicMessage) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *TopicMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *TopicMessage) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_packetbroker_api_routing_v2_service_proto protoreflect.FileDescriptor

var file_packetbroker_api_routing_v2_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x71, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64,
	0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x1c, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x76, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x55, 0x49, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x49, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x1a, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x48, 0x0a,
	0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x6a, 0x0a, 0x12, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x18, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x51, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0c, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x53, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa2, 0x03, 0x0a, 0x06, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x69,
	0x6e, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x39, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x36,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x7e, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x1a, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_routing_v2_service_proto_rawDescOnce sync.Once
	file_packetbroker_api_routing_v2_service_proto_rawDescData = file_packetbroker_api_routing_v2_service_proto_rawDesc
)

func file_packetbroker_api_routing_v2_service_proto_rawDescGZIP() []byte {
	file_packetbroker_api_routing_v2_service_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_routing_v2_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_routing_v2_service_proto_rawDescData)
	})
	return file_packetbroker_api_routing_v2_service_proto_rawDescData
}

var file_packetbroker_api_routing_v2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_packetbroker_api_routing_v2_service_proto_goTypes = []interface{}{
	(*ListUplinkRoutesRequest)(nil),       // 0: org.packetbroker.routing.v2.ListUplinkRoutesRequest
	(*ListUplinkRoutesResponse)(nil),      // 1: org.packetbroker.routing.v2.ListUplinkRoutesResponse
	(*ListJoinRequestRoutesRequest)(nil),  // 2: org.packetbroker.routing.v2.ListJoinRequestRoutesRequest
	(*ListJoinRequestRoutesResponse)(nil), // 3: org.packetbroker.routing.v2.ListJoinRequestRoutesResponse
	(*ListNetworkTargetsRequest)(nil),     // 4: org.packetbroker.routing.v2.ListNetworkTargetsRequest
	(*ListNetworkTargetsResponse)(nil),    // 5: org.packetbroker.routing.v2.ListNetworkTargetsResponse
	(*TopicSubscription)(nil),             // 6: org.packetbroker.routing.v2.TopicSubscription
	(*TopicSubscriptions)(nil),            // 7: org.packetbroker.routing.v2.TopicSubscriptions
	(*TopicSubscriptionsChange)(nil),      // 8: org.packetbroker.routing.v2.TopicSubscriptionsChange
	(*TopicMessage)(nil),                  // 9: org.packetbroker.routing.v2.TopicMessage
	nil,                                   // 10: org.packetbroker.routing.v2.TopicMessage.MetadataEntry
	(*v3.DevAddrPrefixRoute)(nil),         // 11: org.packetbroker.v3.DevAddrPrefixRoute
	(*v3.JoinEUIPrefixRoute)(nil),         // 12: org.packetbroker.v3.JoinEUIPrefixRoute
	(*v3.NetworkTarget)(nil),              // 13: org.packetbroker.v3.NetworkTarget
}
var file_packetbroker_api_routing_v2_service_proto_depIdxs = []int32{
	11, // 0: org.packetbroker.routing.v2.ListUplinkRoutesResponse.routes:type_name -> org.packetbroker.v3.DevAddrPrefixRoute
	12, // 1: org.packetbroker.routing.v2.ListJoinRequestRoutesResponse.routes:type_name -> org.packetbroker.v3.JoinEUIPrefixRoute
	13, // 2: org.packetbroker.routing.v2.ListNetworkTargetsResponse.targets:type_name -> org.packetbroker.v3.NetworkTarget
	6,  // 3: org.packetbroker.routing.v2.TopicSubscriptions.subscriptions:type_name -> org.packetbroker.routing.v2.TopicSubscription
	7,  // 4: org.packetbroker.routing.v2.TopicSubscriptionsChange.subscribes:type_name -> org.packetbroker.routing.v2.TopicSubscriptions
	7,  // 5: org.packetbroker.routing.v2.TopicSubscriptionsChange.unsubscribes:type_name -> org.packetbroker.routing.v2.TopicSubscriptions
	10, // 6: org.packetbroker.routing.v2.TopicMessage.metadata:type_name -> org.packetbroker.routing.v2.TopicMessage.MetadataEntry
	0,  // 7: org.packetbroker.routing.v2.Routes.ListUplinkRoutes:input_type -> org.packetbroker.routing.v2.ListUplinkRoutesRequest
	2,  // 8: org.packetbroker.routing.v2.Routes.ListJoinRequestRoutes:input_type -> org.packetbroker.routing.v2.ListJoinRequestRoutesRequest
	4,  // 9: org.packetbroker.routing.v2.Routes.ListNetworkTargets:input_type -> org.packetbroker.routing.v2.ListNetworkTargetsRequest
	8,  // 10: org.packetbroker.routing.v2.Publisher.Subscribe:input_type -> org.packetbroker.routing.v2.TopicSubscriptionsChange
	1,  // 11: org.packetbroker.routing.v2.Routes.ListUplinkRoutes:output_type -> org.packetbroker.routing.v2.ListUplinkRoutesResponse
	3,  // 12: org.packetbroker.routing.v2.Routes.ListJoinRequestRoutes:output_type -> org.packetbroker.routing.v2.ListJoinRequestRoutesResponse
	5,  // 13: org.packetbroker.routing.v2.Routes.ListNetworkTargets:output_type -> org.packetbroker.routing.v2.ListNetworkTargetsResponse
	9,  // 14: org.packetbroker.routing.v2.Publisher.Subscribe:output_type -> org.packetbroker.routing.v2.TopicMessage
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_packetbroker_api_routing_v2_service_proto_init() }
func file_packetbroker_api_routing_v2_service_proto_init() {
	if File_packetbroker_api_routing_v2_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_routing_v2_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUplinkRoutesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUplinkRoutesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJoinRequestRoutesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJoinRequestRoutesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworkTargetsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworkTargetsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicSubscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicSubscriptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicSubscriptionsChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_packetbroker_api_routing_v2_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_packetbroker_api_routing_v2_service_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*TopicSubscriptionsChange_Subscribes)(nil),
		(*TopicSubscriptionsChange_Unsubscribes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packetbroker_api_routing_v2_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_packetbroker_api_routing_v2_service_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_routing_v2_service_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_routing_v2_service_proto_msgTypes,
	}.Build()
	File_packetbroker_api_routing_v2_service_proto = out.File
	file_packetbroker_api_routing_v2_service_proto_rawDesc = nil
	file_packetbroker_api_routing_v2_service_proto_goTypes = nil
	file_packetbroker_api_routing_v2_service_proto_depIdxs = nil
}
