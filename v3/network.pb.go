// Copyright © 2020 The Things Industries B.V.
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
// source: packetbroker/api/v3/network.proto

package packetbroker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Member.
	NetId uint32 `protobuf:"varint,1,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	// Name of the network.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// DevAddr prefixes of the tenant.
	DevAddrBlocks []*DevAddrBlock `protobuf:"bytes,3,rep,name=dev_addr_blocks,json=devAddrBlocks,proto3" json:"dev_addr_blocks,omitempty"`
	// Administrative contact.
	AdministrativeContact *ContactInfo `protobuf:"bytes,4,opt,name=administrative_contact,json=administrativeContact,proto3" json:"administrative_contact,omitempty"`
	// Technical contact.
	TechnicalContact *ContactInfo `protobuf:"bytes,5,opt,name=technical_contact,json=technicalContact,proto3" json:"technical_contact,omitempty"`
	// Indicates whether the network is publicly listed in the catalog.
	Listed bool `protobuf:"varint,6,opt,name=listed,proto3" json:"listed,omitempty"`
	// Optional target information.
	Target *Target `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetNetId() uint32 {
	if x != nil {
		return x.NetId
	}
	return 0
}

func (x *Network) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Network) GetDevAddrBlocks() []*DevAddrBlock {
	if x != nil {
		return x.DevAddrBlocks
	}
	return nil
}

func (x *Network) GetAdministrativeContact() *ContactInfo {
	if x != nil {
		return x.AdministrativeContact
	}
	return nil
}

func (x *Network) GetTechnicalContact() *ContactInfo {
	if x != nil {
		return x.TechnicalContact
	}
	return nil
}

func (x *Network) GetListed() bool {
	if x != nil {
		return x.Listed
	}
	return false
}

func (x *Network) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Member.
	NetId uint32 `protobuf:"varint,1,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	// ID assigned by the Member.
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Name of the tenant.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// DevAddr prefixes of the tenant.
	DevAddrBlocks []*DevAddrBlock `protobuf:"bytes,5,rep,name=dev_addr_blocks,json=devAddrBlocks,proto3" json:"dev_addr_blocks,omitempty"`
	// Administrative contact.
	AdministrativeContact *ContactInfo `protobuf:"bytes,6,opt,name=administrative_contact,json=administrativeContact,proto3" json:"administrative_contact,omitempty"`
	// Technical contact.
	TechnicalContact *ContactInfo `protobuf:"bytes,7,opt,name=technical_contact,json=technicalContact,proto3" json:"technical_contact,omitempty"`
	// Indicates whether the tenant is listed in the catalog.
	Listed bool `protobuf:"varint,8,opt,name=listed,proto3" json:"listed,omitempty"`
	// Optional target information.
	Target *Target `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{1}
}

func (x *Tenant) GetNetId() uint32 {
	if x != nil {
		return x.NetId
	}
	return 0
}

func (x *Tenant) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetDevAddrBlocks() []*DevAddrBlock {
	if x != nil {
		return x.DevAddrBlocks
	}
	return nil
}

func (x *Tenant) GetAdministrativeContact() *ContactInfo {
	if x != nil {
		return x.AdministrativeContact
	}
	return nil
}

func (x *Tenant) GetTechnicalContact() *ContactInfo {
	if x != nil {
		return x.TechnicalContact
	}
	return nil
}

func (x *Tenant) GetListed() bool {
	if x != nil {
		return x.Listed
	}
	return false
}

func (x *Tenant) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

type NetworkOrTenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*NetworkOrTenant_Network
	//	*NetworkOrTenant_Tenant
	Value isNetworkOrTenant_Value `protobuf_oneof:"value"`
}

func (x *NetworkOrTenant) Reset() {
	*x = NetworkOrTenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkOrTenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkOrTenant) ProtoMessage() {}

func (x *NetworkOrTenant) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkOrTenant.ProtoReflect.Descriptor instead.
func (*NetworkOrTenant) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{2}
}

func (m *NetworkOrTenant) GetValue() isNetworkOrTenant_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *NetworkOrTenant) GetNetwork() *Network {
	if x, ok := x.GetValue().(*NetworkOrTenant_Network); ok {
		return x.Network
	}
	return nil
}

func (x *NetworkOrTenant) GetTenant() *Tenant {
	if x, ok := x.GetValue().(*NetworkOrTenant_Tenant); ok {
		return x.Tenant
	}
	return nil
}

type isNetworkOrTenant_Value interface {
	isNetworkOrTenant_Value()
}

type NetworkOrTenant_Network struct {
	Network *Network `protobuf:"bytes,1,opt,name=network,proto3,oneof"`
}

type NetworkOrTenant_Tenant struct {
	Tenant *Tenant `protobuf:"bytes,2,opt,name=tenant,proto3,oneof"`
}

func (*NetworkOrTenant_Network) isNetworkOrTenant_Value() {}

func (*NetworkOrTenant_Tenant) isNetworkOrTenant_Value() {}

type DevAddrPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// Significant bits of value.
	Length uint32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *DevAddrPrefix) Reset() {
	*x = DevAddrPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevAddrPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevAddrPrefix) ProtoMessage() {}

func (x *DevAddrPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevAddrPrefix.ProtoReflect.Descriptor instead.
func (*DevAddrPrefix) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{3}
}

func (x *DevAddrPrefix) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *DevAddrPrefix) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type DevAddrBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix *DevAddrPrefix `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Cluster of the Home Network Member (optional).
	HomeNetworkClusterId string `protobuf:"bytes,2,opt,name=home_network_cluster_id,json=homeNetworkClusterId,proto3" json:"home_network_cluster_id,omitempty"`
}

func (x *DevAddrBlock) Reset() {
	*x = DevAddrBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevAddrBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevAddrBlock) ProtoMessage() {}

func (x *DevAddrBlock) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevAddrBlock.ProtoReflect.Descriptor instead.
func (*DevAddrBlock) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{4}
}

func (x *DevAddrBlock) GetPrefix() *DevAddrPrefix {
	if x != nil {
		return x.Prefix
	}
	return nil
}

func (x *DevAddrBlock) GetHomeNetworkClusterId() string {
	if x != nil {
		return x.HomeNetworkClusterId
	}
	return ""
}

type NetworkAPIKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the API key.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// LoRa Alliance NetID of the Member.
	NetId uint32 `protobuf:"varint,2,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	// ID assigned by the Member.
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Cluster of the Member.
	ClusterId string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Network rights.
	Rights []Right `protobuf:"varint,7,rep,packed,name=rights,proto3,enum=org.packetbroker.v3.Right" json:"rights,omitempty"`
	// Last authentication timestamp.
	AuthenticatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=authenticated_at,json=authenticatedAt,proto3" json:"authenticated_at,omitempty"`
	// Secret key value.
	Key string `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	// Key state.
	State APIKeyState `protobuf:"varint,8,opt,name=state,proto3,enum=org.packetbroker.v3.APIKeyState" json:"state,omitempty"`
}

func (x *NetworkAPIKey) Reset() {
	*x = NetworkAPIKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAPIKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAPIKey) ProtoMessage() {}

func (x *NetworkAPIKey) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAPIKey.ProtoReflect.Descriptor instead.
func (*NetworkAPIKey) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkAPIKey) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *NetworkAPIKey) GetNetId() uint32 {
	if x != nil {
		return x.NetId
	}
	return 0
}

func (x *NetworkAPIKey) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *NetworkAPIKey) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *NetworkAPIKey) GetRights() []Right {
	if x != nil {
		return x.Rights
	}
	return nil
}

func (x *NetworkAPIKey) GetAuthenticatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AuthenticatedAt
	}
	return nil
}

func (x *NetworkAPIKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NetworkAPIKey) GetState() APIKeyState {
	if x != nil {
		return x.State
	}
	return APIKeyState_REQUESTED
}

type NetworkTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoRa Alliance NetID of the Member.
	NetId uint32 `protobuf:"varint,1,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	// Tenant ID assigned by the Member (optional).
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Target information.
	Target *Target `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *NetworkTarget) Reset() {
	*x = NetworkTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkTarget) ProtoMessage() {}

func (x *NetworkTarget) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkTarget.ProtoReflect.Descriptor instead.
func (*NetworkTarget) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_network_proto_rawDescGZIP(), []int{6}
}

func (x *NetworkTarget) GetNetId() uint32 {
	if x != nil {
		return x.NetId
	}
	return 0
}

func (x *NetworkTarget) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *NetworkTarget) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

var File_packetbroker_api_v3_network_proto protoreflect.FileDescriptor

var file_packetbroker_api_v3_network_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x33, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf4, 0x02, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x6e,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x57, 0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x4d, 0x0a, 0x11, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x33, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x96, 0x03, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x65, 0x76,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x57, 0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x4d, 0x0a,
	0x11, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22,
	0x8b, 0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x72, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x35, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x81, 0x01, 0x0a,
	0x0c, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3a, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x35, 0x0a, 0x17, 0x68, 0x6f, 0x6d,
	0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x68, 0x6f, 0x6d, 0x65,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xbe, 0x02, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x12, 0x45, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x78, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_v3_network_proto_rawDescOnce sync.Once
	file_packetbroker_api_v3_network_proto_rawDescData = file_packetbroker_api_v3_network_proto_rawDesc
)

func file_packetbroker_api_v3_network_proto_rawDescGZIP() []byte {
	file_packetbroker_api_v3_network_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_v3_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_v3_network_proto_rawDescData)
	})
	return file_packetbroker_api_v3_network_proto_rawDescData
}

var file_packetbroker_api_v3_network_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_packetbroker_api_v3_network_proto_goTypes = []interface{}{
	(*Network)(nil),               // 0: org.packetbroker.v3.Network
	(*Tenant)(nil),                // 1: org.packetbroker.v3.Tenant
	(*NetworkOrTenant)(nil),       // 2: org.packetbroker.v3.NetworkOrTenant
	(*DevAddrPrefix)(nil),         // 3: org.packetbroker.v3.DevAddrPrefix
	(*DevAddrBlock)(nil),          // 4: org.packetbroker.v3.DevAddrBlock
	(*NetworkAPIKey)(nil),         // 5: org.packetbroker.v3.NetworkAPIKey
	(*NetworkTarget)(nil),         // 6: org.packetbroker.v3.NetworkTarget
	(*ContactInfo)(nil),           // 7: org.packetbroker.v3.ContactInfo
	(*Target)(nil),                // 8: org.packetbroker.v3.Target
	(Right)(0),                    // 9: org.packetbroker.v3.Right
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(APIKeyState)(0),              // 11: org.packetbroker.v3.APIKeyState
}
var file_packetbroker_api_v3_network_proto_depIdxs = []int32{
	4,  // 0: org.packetbroker.v3.Network.dev_addr_blocks:type_name -> org.packetbroker.v3.DevAddrBlock
	7,  // 1: org.packetbroker.v3.Network.administrative_contact:type_name -> org.packetbroker.v3.ContactInfo
	7,  // 2: org.packetbroker.v3.Network.technical_contact:type_name -> org.packetbroker.v3.ContactInfo
	8,  // 3: org.packetbroker.v3.Network.target:type_name -> org.packetbroker.v3.Target
	4,  // 4: org.packetbroker.v3.Tenant.dev_addr_blocks:type_name -> org.packetbroker.v3.DevAddrBlock
	7,  // 5: org.packetbroker.v3.Tenant.administrative_contact:type_name -> org.packetbroker.v3.ContactInfo
	7,  // 6: org.packetbroker.v3.Tenant.technical_contact:type_name -> org.packetbroker.v3.ContactInfo
	8,  // 7: org.packetbroker.v3.Tenant.target:type_name -> org.packetbroker.v3.Target
	0,  // 8: org.packetbroker.v3.NetworkOrTenant.network:type_name -> org.packetbroker.v3.Network
	1,  // 9: org.packetbroker.v3.NetworkOrTenant.tenant:type_name -> org.packetbroker.v3.Tenant
	3,  // 10: org.packetbroker.v3.DevAddrBlock.prefix:type_name -> org.packetbroker.v3.DevAddrPrefix
	9,  // 11: org.packetbroker.v3.NetworkAPIKey.rights:type_name -> org.packetbroker.v3.Right
	10, // 12: org.packetbroker.v3.NetworkAPIKey.authenticated_at:type_name -> google.protobuf.Timestamp
	11, // 13: org.packetbroker.v3.NetworkAPIKey.state:type_name -> org.packetbroker.v3.APIKeyState
	8,  // 14: org.packetbroker.v3.NetworkTarget.target:type_name -> org.packetbroker.v3.Target
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_packetbroker_api_v3_network_proto_init() }
func file_packetbroker_api_v3_network_proto_init() {
	if File_packetbroker_api_v3_network_proto != nil {
		return
	}
	file_packetbroker_api_v3_contact_proto_init()
	file_packetbroker_api_v3_enums_proto_init()
	file_packetbroker_api_v3_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_v3_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkOrTenant); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevAddrPrefix); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevAddrBlock); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAPIKey); i {
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
		file_packetbroker_api_v3_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkTarget); i {
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
	file_packetbroker_api_v3_network_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*NetworkOrTenant_Network)(nil),
		(*NetworkOrTenant_Tenant)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packetbroker_api_v3_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packetbroker_api_v3_network_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_v3_network_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_v3_network_proto_msgTypes,
	}.Build()
	File_packetbroker_api_v3_network_proto = out.File
	file_packetbroker_api_v3_network_proto_rawDesc = nil
	file_packetbroker_api_v3_network_proto_goTypes = nil
	file_packetbroker_api_v3_network_proto_depIdxs = nil
}
