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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: packetbroker/api/mapping/v2/service.proto

package mappingpb

import (
	v3 "go.packetbroker.org/api/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDefaultGatewayVisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Forwarder Member.
	ForwarderNetId uint32 `protobuf:"varint,1,opt,name=forwarder_net_id,json=forwarderNetId,proto3" json:"forwarder_net_id,omitempty"`
	// Tenant ID managed by the Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,2,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
}

func (x *GetDefaultGatewayVisibilityRequest) Reset() {
	*x = GetDefaultGatewayVisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultGatewayVisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultGatewayVisibilityRequest) ProtoMessage() {}

func (x *GetDefaultGatewayVisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultGatewayVisibilityRequest.ProtoReflect.Descriptor instead.
func (*GetDefaultGatewayVisibilityRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDefaultGatewayVisibilityRequest) GetForwarderNetId() uint32 {
	if x != nil {
		return x.ForwarderNetId
	}
	return 0
}

func (x *GetDefaultGatewayVisibilityRequest) GetForwarderTenantId() string {
	if x != nil {
		return x.ForwarderTenantId
	}
	return ""
}

type GetHomeNetworkGatewayVisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Forwarder Member.
	ForwarderNetId uint32 `protobuf:"varint,1,opt,name=forwarder_net_id,json=forwarderNetId,proto3" json:"forwarder_net_id,omitempty"`
	// Tenant ID managed by the Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,2,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
	// LoRa Alliance NetID of the Home Network Member.
	HomeNetworkNetId uint32 `protobuf:"varint,3,opt,name=home_network_net_id,json=homeNetworkNetId,proto3" json:"home_network_net_id,omitempty"`
	// Tenant ID managed by the Home Network Member.
	HomeNetworkTenantId string `protobuf:"bytes,4,opt,name=home_network_tenant_id,json=homeNetworkTenantId,proto3" json:"home_network_tenant_id,omitempty"`
}

func (x *GetHomeNetworkGatewayVisibilityRequest) Reset() {
	*x = GetHomeNetworkGatewayVisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeNetworkGatewayVisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeNetworkGatewayVisibilityRequest) ProtoMessage() {}

func (x *GetHomeNetworkGatewayVisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeNetworkGatewayVisibilityRequest.ProtoReflect.Descriptor instead.
func (*GetHomeNetworkGatewayVisibilityRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHomeNetworkGatewayVisibilityRequest) GetForwarderNetId() uint32 {
	if x != nil {
		return x.ForwarderNetId
	}
	return 0
}

func (x *GetHomeNetworkGatewayVisibilityRequest) GetForwarderTenantId() string {
	if x != nil {
		return x.ForwarderTenantId
	}
	return ""
}

func (x *GetHomeNetworkGatewayVisibilityRequest) GetHomeNetworkNetId() uint32 {
	if x != nil {
		return x.HomeNetworkNetId
	}
	return 0
}

func (x *GetHomeNetworkGatewayVisibilityRequest) GetHomeNetworkTenantId() string {
	if x != nil {
		return x.HomeNetworkTenantId
	}
	return ""
}

type SetGatewayVisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility *v3.GatewayVisibility `protobuf:"bytes,3,opt,name=visibility,proto3" json:"visibility,omitempty"`
}

func (x *SetGatewayVisibilityRequest) Reset() {
	*x = SetGatewayVisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGatewayVisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGatewayVisibilityRequest) ProtoMessage() {}

func (x *SetGatewayVisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGatewayVisibilityRequest.ProtoReflect.Descriptor instead.
func (*SetGatewayVisibilityRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetGatewayVisibilityRequest) GetVisibility() *v3.GatewayVisibility {
	if x != nil {
		return x.Visibility
	}
	return nil
}

type GetGatewayVisibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility *v3.GatewayVisibility `protobuf:"bytes,1,opt,name=visibility,proto3" json:"visibility,omitempty"`
}

func (x *GetGatewayVisibilityResponse) Reset() {
	*x = GetGatewayVisibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayVisibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayVisibilityResponse) ProtoMessage() {}

func (x *GetGatewayVisibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayVisibilityResponse.ProtoReflect.Descriptor instead.
func (*GetGatewayVisibilityResponse) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetGatewayVisibilityResponse) GetVisibility() *v3.GatewayVisibility {
	if x != nil {
		return x.Visibility
	}
	return nil
}

type UpdateGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Forwarder Member.
	ForwarderNetId uint32 `protobuf:"varint,1,opt,name=forwarder_net_id,json=forwarderNetId,proto3" json:"forwarder_net_id,omitempty"`
	// Tenant ID managed by the Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,2,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
	// Forwarder cluster of the Forwarder Member (optional).
	ForwarderClusterId string `protobuf:"bytes,3,opt,name=forwarder_cluster_id,json=forwarderClusterId,proto3" json:"forwarder_cluster_id,omitempty"`
	// Identifier of the gateway.
	ForwarderGatewayId *v3.GatewayIdentifier `protobuf:"bytes,4,opt,name=forwarder_gateway_id,json=forwarderGatewayId,proto3" json:"forwarder_gateway_id,omitempty"`
	// Gateway location.
	// This field gets updated when a value is set.
	GatewayLocation *v3.GatewayLocationValue `protobuf:"bytes,5,opt,name=gateway_location,json=gatewayLocation,proto3" json:"gateway_location,omitempty"`
	// Administrative contact.
	// This field gets updated when a value is set.
	AdministrativeContact *v3.ContactInfoValue `protobuf:"bytes,6,opt,name=administrative_contact,json=administrativeContact,proto3" json:"administrative_contact,omitempty"`
	// Technical contact.
	// This field gets updated when a value is set.
	TechnicalContact *v3.ContactInfoValue `protobuf:"bytes,7,opt,name=technical_contact,json=technicalContact,proto3" json:"technical_contact,omitempty"`
	// Indicates whether the gateway is online.
	// This field gets updated when a value is set.
	Online *wrapperspb.BoolValue `protobuf:"bytes,8,opt,name=online,proto3" json:"online,omitempty"`
	// If the gateway is online, this value indicates the time-to-live for the online status. This value must be set when
	// the online field is set and true.
	// When the online status expires, the gateway will be offline.
	// To keep the gateway online, update online status before the status expires.
	OnlineTtl *durationpb.Duration `protobuf:"bytes,9,opt,name=online_ttl,json=onlineTtl,proto3" json:"online_ttl,omitempty"`
	// Frequency plan of the gateway.
	// This field gets updated when a value is set.
	FrequencyPlan *v3.GatewayFrequencyPlan `protobuf:"bytes,10,opt,name=frequency_plan,json=frequencyPlan,proto3" json:"frequency_plan,omitempty"`
	// Received packets rate (number of packets per hour).
	// This field gets updated when a value is set.
	RxRate *wrapperspb.FloatValue `protobuf:"bytes,11,opt,name=rx_rate,json=rxRate,proto3" json:"rx_rate,omitempty"`
	// Transmitted packets rate (number of packets per hour).
	// This field gets updated when a value is set.
	TxRate *wrapperspb.FloatValue `protobuf:"bytes,12,opt,name=tx_rate,json=txRate,proto3" json:"tx_rate,omitempty"`
}

func (x *UpdateGatewayRequest) Reset() {
	*x = UpdateGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGatewayRequest) ProtoMessage() {}

func (x *UpdateGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_mapping_v2_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGatewayRequest.ProtoReflect.Descriptor instead.
func (*UpdateGatewayRequest) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGatewayRequest) GetForwarderNetId() uint32 {
	if x != nil {
		return x.ForwarderNetId
	}
	return 0
}

func (x *UpdateGatewayRequest) GetForwarderTenantId() string {
	if x != nil {
		return x.ForwarderTenantId
	}
	return ""
}

func (x *UpdateGatewayRequest) GetForwarderClusterId() string {
	if x != nil {
		return x.ForwarderClusterId
	}
	return ""
}

func (x *UpdateGatewayRequest) GetForwarderGatewayId() *v3.GatewayIdentifier {
	if x != nil {
		return x.ForwarderGatewayId
	}
	return nil
}

func (x *UpdateGatewayRequest) GetGatewayLocation() *v3.GatewayLocationValue {
	if x != nil {
		return x.GatewayLocation
	}
	return nil
}

func (x *UpdateGatewayRequest) GetAdministrativeContact() *v3.ContactInfoValue {
	if x != nil {
		return x.AdministrativeContact
	}
	return nil
}

func (x *UpdateGatewayRequest) GetTechnicalContact() *v3.ContactInfoValue {
	if x != nil {
		return x.TechnicalContact
	}
	return nil
}

func (x *UpdateGatewayRequest) GetOnline() *wrapperspb.BoolValue {
	if x != nil {
		return x.Online
	}
	return nil
}

func (x *UpdateGatewayRequest) GetOnlineTtl() *durationpb.Duration {
	if x != nil {
		return x.OnlineTtl
	}
	return nil
}

func (x *UpdateGatewayRequest) GetFrequencyPlan() *v3.GatewayFrequencyPlan {
	if x != nil {
		return x.FrequencyPlan
	}
	return nil
}

func (x *UpdateGatewayRequest) GetRxRate() *wrapperspb.FloatValue {
	if x != nil {
		return x.RxRate
	}
	return nil
}

func (x *UpdateGatewayRequest) GetTxRate() *wrapperspb.FloatValue {
	if x != nil {
		return x.TxRate
	}
	return nil
}

var File_packetbroker_api_mapping_v2_service_proto protoreflect.FileDescriptor

var file_packetbroker_api_mapping_v2_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x22, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x26,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x13, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x68,
	0x6f, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x16, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x66, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0xb0, 0x06, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x14, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x12,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x54, 0x0a, 0x10, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x52, 0x0a, 0x11, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38,
	0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x74, 0x6c, 0x12, 0x50, 0x0a, 0x0e, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x46, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x0d, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x78,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x78, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x74, 0x78, 0x52, 0x61, 0x74, 0x65, 0x32, 0xa4, 0x04, 0x0a, 0x18, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x92, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x38, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x9a, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x43, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6c, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x38, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x64, 0x0a,
	0x06, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_mapping_v2_service_proto_rawDescOnce sync.Once
	file_packetbroker_api_mapping_v2_service_proto_rawDescData = file_packetbroker_api_mapping_v2_service_proto_rawDesc
)

func file_packetbroker_api_mapping_v2_service_proto_rawDescGZIP() []byte {
	file_packetbroker_api_mapping_v2_service_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_mapping_v2_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_mapping_v2_service_proto_rawDescData)
	})
	return file_packetbroker_api_mapping_v2_service_proto_rawDescData
}

var file_packetbroker_api_mapping_v2_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_packetbroker_api_mapping_v2_service_proto_goTypes = []interface{}{
	(*GetDefaultGatewayVisibilityRequest)(nil),     // 0: org.packetbroker.mapping.v2.GetDefaultGatewayVisibilityRequest
	(*GetHomeNetworkGatewayVisibilityRequest)(nil), // 1: org.packetbroker.mapping.v2.GetHomeNetworkGatewayVisibilityRequest
	(*SetGatewayVisibilityRequest)(nil),            // 2: org.packetbroker.mapping.v2.SetGatewayVisibilityRequest
	(*GetGatewayVisibilityResponse)(nil),           // 3: org.packetbroker.mapping.v2.GetGatewayVisibilityResponse
	(*UpdateGatewayRequest)(nil),                   // 4: org.packetbroker.mapping.v2.UpdateGatewayRequest
	(*v3.GatewayVisibility)(nil),                   // 5: org.packetbroker.v3.GatewayVisibility
	(*v3.GatewayIdentifier)(nil),                   // 6: org.packetbroker.v3.GatewayIdentifier
	(*v3.GatewayLocationValue)(nil),                // 7: org.packetbroker.v3.GatewayLocationValue
	(*v3.ContactInfoValue)(nil),                    // 8: org.packetbroker.v3.ContactInfoValue
	(*wrapperspb.BoolValue)(nil),                   // 9: google.protobuf.BoolValue
	(*durationpb.Duration)(nil),                    // 10: google.protobuf.Duration
	(*v3.GatewayFrequencyPlan)(nil),                // 11: org.packetbroker.v3.GatewayFrequencyPlan
	(*wrapperspb.FloatValue)(nil),                  // 12: google.protobuf.FloatValue
	(*emptypb.Empty)(nil),                          // 13: google.protobuf.Empty
}
var file_packetbroker_api_mapping_v2_service_proto_depIdxs = []int32{
	5,  // 0: org.packetbroker.mapping.v2.SetGatewayVisibilityRequest.visibility:type_name -> org.packetbroker.v3.GatewayVisibility
	5,  // 1: org.packetbroker.mapping.v2.GetGatewayVisibilityResponse.visibility:type_name -> org.packetbroker.v3.GatewayVisibility
	6,  // 2: org.packetbroker.mapping.v2.UpdateGatewayRequest.forwarder_gateway_id:type_name -> org.packetbroker.v3.GatewayIdentifier
	7,  // 3: org.packetbroker.mapping.v2.UpdateGatewayRequest.gateway_location:type_name -> org.packetbroker.v3.GatewayLocationValue
	8,  // 4: org.packetbroker.mapping.v2.UpdateGatewayRequest.administrative_contact:type_name -> org.packetbroker.v3.ContactInfoValue
	8,  // 5: org.packetbroker.mapping.v2.UpdateGatewayRequest.technical_contact:type_name -> org.packetbroker.v3.ContactInfoValue
	9,  // 6: org.packetbroker.mapping.v2.UpdateGatewayRequest.online:type_name -> google.protobuf.BoolValue
	10, // 7: org.packetbroker.mapping.v2.UpdateGatewayRequest.online_ttl:type_name -> google.protobuf.Duration
	11, // 8: org.packetbroker.mapping.v2.UpdateGatewayRequest.frequency_plan:type_name -> org.packetbroker.v3.GatewayFrequencyPlan
	12, // 9: org.packetbroker.mapping.v2.UpdateGatewayRequest.rx_rate:type_name -> google.protobuf.FloatValue
	12, // 10: org.packetbroker.mapping.v2.UpdateGatewayRequest.tx_rate:type_name -> google.protobuf.FloatValue
	0,  // 11: org.packetbroker.mapping.v2.GatewayVisibilityManager.GetDefaultVisibility:input_type -> org.packetbroker.mapping.v2.GetDefaultGatewayVisibilityRequest
	2,  // 12: org.packetbroker.mapping.v2.GatewayVisibilityManager.SetDefaultVisibility:input_type -> org.packetbroker.mapping.v2.SetGatewayVisibilityRequest
	1,  // 13: org.packetbroker.mapping.v2.GatewayVisibilityManager.GetHomeNetworkVisibility:input_type -> org.packetbroker.mapping.v2.GetHomeNetworkGatewayVisibilityRequest
	2,  // 14: org.packetbroker.mapping.v2.GatewayVisibilityManager.SetHomeNetworkVisibility:input_type -> org.packetbroker.mapping.v2.SetGatewayVisibilityRequest
	4,  // 15: org.packetbroker.mapping.v2.Mapper.UpdateGateway:input_type -> org.packetbroker.mapping.v2.UpdateGatewayRequest
	3,  // 16: org.packetbroker.mapping.v2.GatewayVisibilityManager.GetDefaultVisibility:output_type -> org.packetbroker.mapping.v2.GetGatewayVisibilityResponse
	13, // 17: org.packetbroker.mapping.v2.GatewayVisibilityManager.SetDefaultVisibility:output_type -> google.protobuf.Empty
	3,  // 18: org.packetbroker.mapping.v2.GatewayVisibilityManager.GetHomeNetworkVisibility:output_type -> org.packetbroker.mapping.v2.GetGatewayVisibilityResponse
	13, // 19: org.packetbroker.mapping.v2.GatewayVisibilityManager.SetHomeNetworkVisibility:output_type -> google.protobuf.Empty
	13, // 20: org.packetbroker.mapping.v2.Mapper.UpdateGateway:output_type -> google.protobuf.Empty
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_packetbroker_api_mapping_v2_service_proto_init() }
func file_packetbroker_api_mapping_v2_service_proto_init() {
	if File_packetbroker_api_mapping_v2_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_mapping_v2_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultGatewayVisibilityRequest); i {
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
		file_packetbroker_api_mapping_v2_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeNetworkGatewayVisibilityRequest); i {
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
		file_packetbroker_api_mapping_v2_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGatewayVisibilityRequest); i {
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
		file_packetbroker_api_mapping_v2_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGatewayVisibilityResponse); i {
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
		file_packetbroker_api_mapping_v2_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGatewayRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packetbroker_api_mapping_v2_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_packetbroker_api_mapping_v2_service_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_mapping_v2_service_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_mapping_v2_service_proto_msgTypes,
	}.Build()
	File_packetbroker_api_mapping_v2_service_proto = out.File
	file_packetbroker_api_mapping_v2_service_proto_rawDesc = nil
	file_packetbroker_api_mapping_v2_service_proto_goTypes = nil
	file_packetbroker_api_mapping_v2_service_proto_depIdxs = nil
}
