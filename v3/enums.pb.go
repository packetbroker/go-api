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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: packetbroker/api/v3/enums.proto

package packetbroker

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// LoRaWAN region as defined in Regional Parameters.
type Region int32

const (
	Region_UNKNOWN_REGION Region = 0
	Region_EU_863_870     Region = 1
	Region_US_902_928     Region = 2
	Region_CN_779_787     Region = 3
	Region_EU_433         Region = 4
	Region_AU_915_928     Region = 5
	Region_CN_470_510     Region = 6
	Region_AS_923         Region = 7
	// AS923 group 1
	Region_AS_923_1 Region = 11
	// AS923 group 2
	Region_AS_923_2 Region = 12
	// AS923 group 3
	Region_AS_923_3   Region = 13
	Region_KR_920_923 Region = 8
	Region_IN_865_867 Region = 9
	Region_RU_864_870 Region = 10
	// LoRa 2.4 GHz
	Region_WW2G4 Region = 14
)

// Enum value maps for Region.
var (
	Region_name = map[int32]string{
		0:  "UNKNOWN_REGION",
		1:  "EU_863_870",
		2:  "US_902_928",
		3:  "CN_779_787",
		4:  "EU_433",
		5:  "AU_915_928",
		6:  "CN_470_510",
		7:  "AS_923",
		11: "AS_923_1",
		12: "AS_923_2",
		13: "AS_923_3",
		8:  "KR_920_923",
		9:  "IN_865_867",
		10: "RU_864_870",
		14: "WW2G4",
	}
	Region_value = map[string]int32{
		"UNKNOWN_REGION": 0,
		"EU_863_870":     1,
		"US_902_928":     2,
		"CN_779_787":     3,
		"EU_433":         4,
		"AU_915_928":     5,
		"CN_470_510":     6,
		"AS_923":         7,
		"AS_923_1":       11,
		"AS_923_2":       12,
		"AS_923_3":       13,
		"KR_920_923":     8,
		"IN_865_867":     9,
		"RU_864_870":     10,
		"WW2G4":          14,
	}
)

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}

func (x Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Region) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[0].Descriptor()
}

func (Region) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[0]
}

func (x Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region.Descriptor instead.
func (Region) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{0}
}

// LoRaWAN class for transmitting downlink.
type DownlinkMessageClass int32

const (
	DownlinkMessageClass_CLASS_A DownlinkMessageClass = 0
	DownlinkMessageClass_CLASS_B DownlinkMessageClass = 1
	DownlinkMessageClass_CLASS_C DownlinkMessageClass = 2
)

// Enum value maps for DownlinkMessageClass.
var (
	DownlinkMessageClass_name = map[int32]string{
		0: "CLASS_A",
		1: "CLASS_B",
		2: "CLASS_C",
	}
	DownlinkMessageClass_value = map[string]int32{
		"CLASS_A": 0,
		"CLASS_B": 1,
		"CLASS_C": 2,
	}
)

func (x DownlinkMessageClass) Enum() *DownlinkMessageClass {
	p := new(DownlinkMessageClass)
	*p = x
	return p
}

func (x DownlinkMessageClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownlinkMessageClass) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[1].Descriptor()
}

func (DownlinkMessageClass) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[1]
}

func (x DownlinkMessageClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownlinkMessageClass.Descriptor instead.
func (DownlinkMessageClass) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{1}
}

// Priority of a downlink message.
type DownlinkMessagePriority int32

const (
	// Lowest priority.
	DownlinkMessagePriority_LOWEST DownlinkMessagePriority = 0
	// Low priority.
	DownlinkMessagePriority_LOW DownlinkMessagePriority = 1
	// Normal priority.
	DownlinkMessagePriority_NORMAL DownlinkMessagePriority = 2
	// High priority.
	// Sets the HiPriority flag in the LoRaWAN Backend Interfaces XmitDataReq message.
	DownlinkMessagePriority_HIGH DownlinkMessagePriority = 3
	// Highest priority.
	// Sets the HiPriority flag in the LoRaWAN Backend Interfaces XmitDataReq message.
	DownlinkMessagePriority_HIGHEST DownlinkMessagePriority = 4
)

// Enum value maps for DownlinkMessagePriority.
var (
	DownlinkMessagePriority_name = map[int32]string{
		0: "LOWEST",
		1: "LOW",
		2: "NORMAL",
		3: "HIGH",
		4: "HIGHEST",
	}
	DownlinkMessagePriority_value = map[string]int32{
		"LOWEST":  0,
		"LOW":     1,
		"NORMAL":  2,
		"HIGH":    3,
		"HIGHEST": 4,
	}
)

func (x DownlinkMessagePriority) Enum() *DownlinkMessagePriority {
	p := new(DownlinkMessagePriority)
	*p = x
	return p
}

func (x DownlinkMessagePriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownlinkMessagePriority) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[2].Descriptor()
}

func (DownlinkMessagePriority) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[2]
}

func (x DownlinkMessagePriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownlinkMessagePriority.Descriptor instead.
func (DownlinkMessagePriority) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{2}
}

type MessageDeliveryState int32

const (
	// The message has been accepted for forwarding or delivery.
	MessageDeliveryState_ACCEPTED MessageDeliveryState = 0
	// The message has been forwarded to another host for delivery.
	MessageDeliveryState_FORWARDED MessageDeliveryState = 1
	// The message has been delivered to the receiver.
	MessageDeliveryState_DELIVERED MessageDeliveryState = 2
	// The message has been processed by the receiver.
	MessageDeliveryState_PROCESSED MessageDeliveryState = 3
)

// Enum value maps for MessageDeliveryState.
var (
	MessageDeliveryState_name = map[int32]string{
		0: "ACCEPTED",
		1: "FORWARDED",
		2: "DELIVERED",
		3: "PROCESSED",
	}
	MessageDeliveryState_value = map[string]int32{
		"ACCEPTED":  0,
		"FORWARDED": 1,
		"DELIVERED": 2,
		"PROCESSED": 3,
	}
)

func (x MessageDeliveryState) Enum() *MessageDeliveryState {
	p := new(MessageDeliveryState)
	*p = x
	return p
}

func (x MessageDeliveryState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageDeliveryState) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[3].Descriptor()
}

func (MessageDeliveryState) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[3]
}

func (x MessageDeliveryState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageDeliveryState.Descriptor instead.
func (MessageDeliveryState) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{3}
}

type TerrestrialAntennaPlacement int32

const (
	// Unknown antenna placement.
	TerrestrialAntennaPlacement_UNKNOWN_PLACEMENT TerrestrialAntennaPlacement = 0
	// Indoor antenna placement.
	TerrestrialAntennaPlacement_INDOOR TerrestrialAntennaPlacement = 1
	// Outdoor antenna placement.
	TerrestrialAntennaPlacement_OUTDOOR TerrestrialAntennaPlacement = 2
)

// Enum value maps for TerrestrialAntennaPlacement.
var (
	TerrestrialAntennaPlacement_name = map[int32]string{
		0: "UNKNOWN_PLACEMENT",
		1: "INDOOR",
		2: "OUTDOOR",
	}
	TerrestrialAntennaPlacement_value = map[string]int32{
		"UNKNOWN_PLACEMENT": 0,
		"INDOOR":            1,
		"OUTDOOR":           2,
	}
)

func (x TerrestrialAntennaPlacement) Enum() *TerrestrialAntennaPlacement {
	p := new(TerrestrialAntennaPlacement)
	*p = x
	return p
}

func (x TerrestrialAntennaPlacement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerrestrialAntennaPlacement) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[4].Descriptor()
}

func (TerrestrialAntennaPlacement) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[4]
}

func (x TerrestrialAntennaPlacement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerrestrialAntennaPlacement.Descriptor instead.
func (TerrestrialAntennaPlacement) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{4}
}

type Modulation int32

const (
	// LoRa modulation.
	Modulation_LORA Modulation = 0
	// FSK modulation.
	Modulation_FSK Modulation = 1
)

// Enum value maps for Modulation.
var (
	Modulation_name = map[int32]string{
		0: "LORA",
		1: "FSK",
	}
	Modulation_value = map[string]int32{
		"LORA": 0,
		"FSK":  1,
	}
)

func (x Modulation) Enum() *Modulation {
	p := new(Modulation)
	*p = x
	return p
}

func (x Modulation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Modulation) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[5].Descriptor()
}

func (Modulation) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[5]
}

func (x Modulation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Modulation.Descriptor instead.
func (Modulation) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{5}
}

type Right int32

const (
	// Right to list networks and get network information.
	Right_READ_NETWORK Right = 0
	// Right to read contact information of networks.
	Right_READ_NETWORK_CONTACT Right = 1
	// Right to list network tenants and read tenants information.
	Right_READ_TENANT Right = 2
	// Right to read contact information of tenants.
	Right_READ_TENANT_CONTACT Right = 3
	// Right to list routing policies.
	Right_READ_ROUTING_POLICY Right = 4
	// Right to list the route table.
	Right_READ_ROUTE_TABLE Right = 5
)

// Enum value maps for Right.
var (
	Right_name = map[int32]string{
		0: "READ_NETWORK",
		1: "READ_NETWORK_CONTACT",
		2: "READ_TENANT",
		3: "READ_TENANT_CONTACT",
		4: "READ_ROUTING_POLICY",
		5: "READ_ROUTE_TABLE",
	}
	Right_value = map[string]int32{
		"READ_NETWORK":         0,
		"READ_NETWORK_CONTACT": 1,
		"READ_TENANT":          2,
		"READ_TENANT_CONTACT":  3,
		"READ_ROUTING_POLICY":  4,
		"READ_ROUTE_TABLE":     5,
	}
)

func (x Right) Enum() *Right {
	p := new(Right)
	*p = x
	return p
}

func (x Right) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Right) Descriptor() protoreflect.EnumDescriptor {
	return file_packetbroker_api_v3_enums_proto_enumTypes[6].Descriptor()
}

func (Right) Type() protoreflect.EnumType {
	return &file_packetbroker_api_v3_enums_proto_enumTypes[6]
}

func (x Right) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Right.Descriptor instead.
func (Right) EnumDescriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_enums_proto_rawDescGZIP(), []int{6}
}

var File_packetbroker_api_v3_enums_proto protoreflect.FileDescriptor

var file_packetbroker_api_v3_enums_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2a, 0xe9, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x55, 0x5f, 0x38, 0x36, 0x33, 0x5f,
	0x38, 0x37, 0x30, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x5f, 0x39, 0x30, 0x32, 0x5f,
	0x39, 0x32, 0x38, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4e, 0x5f, 0x37, 0x37, 0x39, 0x5f,
	0x37, 0x38, 0x37, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x55, 0x5f, 0x34, 0x33, 0x33, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x5f, 0x39, 0x31, 0x35, 0x5f, 0x39, 0x32, 0x38, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4e, 0x5f, 0x34, 0x37, 0x30, 0x5f, 0x35, 0x31, 0x30, 0x10,
	0x06, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x53, 0x5f, 0x39, 0x32, 0x33, 0x10, 0x07, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x53, 0x5f, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x53, 0x5f, 0x39, 0x32, 0x33, 0x5f, 0x32, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x5f,
	0x39, 0x32, 0x33, 0x5f, 0x33, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x52, 0x5f, 0x39, 0x32,
	0x30, 0x5f, 0x39, 0x32, 0x33, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x38, 0x36,
	0x35, 0x5f, 0x38, 0x36, 0x37, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x55, 0x5f, 0x38, 0x36,
	0x34, 0x5f, 0x38, 0x37, 0x30, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x57, 0x32, 0x47, 0x34,
	0x10, 0x0e, 0x2a, 0x3d, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x41, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x5f, 0x42, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x43, 0x10,
	0x02, 0x2a, 0x51, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06,
	0x4c, 0x4f, 0x57, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x49, 0x47, 0x48, 0x45,
	0x53, 0x54, 0x10, 0x04, 0x2a, 0x51, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f,
	0x52, 0x57, 0x41, 0x52, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c,
	0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x4d, 0x0a, 0x1b, 0x54, 0x65, 0x72, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x49, 0x4e, 0x44, 0x4f, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x55, 0x54,
	0x44, 0x4f, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x1f, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x52, 0x41, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x46, 0x53, 0x4b, 0x10, 0x01, 0x2a, 0x8c, 0x01, 0x0a, 0x05, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4e, 0x45, 0x54, 0x57,
	0x4f, 0x52, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x41, 0x43, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x04,
	0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_v3_enums_proto_rawDescOnce sync.Once
	file_packetbroker_api_v3_enums_proto_rawDescData = file_packetbroker_api_v3_enums_proto_rawDesc
)

func file_packetbroker_api_v3_enums_proto_rawDescGZIP() []byte {
	file_packetbroker_api_v3_enums_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_v3_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_v3_enums_proto_rawDescData)
	})
	return file_packetbroker_api_v3_enums_proto_rawDescData
}

var file_packetbroker_api_v3_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_packetbroker_api_v3_enums_proto_goTypes = []interface{}{
	(Region)(0),                      // 0: org.packetbroker.v3.Region
	(DownlinkMessageClass)(0),        // 1: org.packetbroker.v3.DownlinkMessageClass
	(DownlinkMessagePriority)(0),     // 2: org.packetbroker.v3.DownlinkMessagePriority
	(MessageDeliveryState)(0),        // 3: org.packetbroker.v3.MessageDeliveryState
	(TerrestrialAntennaPlacement)(0), // 4: org.packetbroker.v3.TerrestrialAntennaPlacement
	(Modulation)(0),                  // 5: org.packetbroker.v3.Modulation
	(Right)(0),                       // 6: org.packetbroker.v3.Right
}
var file_packetbroker_api_v3_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_packetbroker_api_v3_enums_proto_init() }
func file_packetbroker_api_v3_enums_proto_init() {
	if File_packetbroker_api_v3_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packetbroker_api_v3_enums_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packetbroker_api_v3_enums_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_v3_enums_proto_depIdxs,
		EnumInfos:         file_packetbroker_api_v3_enums_proto_enumTypes,
	}.Build()
	File_packetbroker_api_v3_enums_proto = out.File
	file_packetbroker_api_v3_enums_proto_rawDesc = nil
	file_packetbroker_api_v3_enums_proto_goTypes = nil
	file_packetbroker_api_v3_enums_proto_depIdxs = nil
}
