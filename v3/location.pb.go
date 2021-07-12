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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: packetbroker/api/v3/location.proto

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The North–South position (degrees; -90 to +90), where 0 is the equator, North pole is positive, South pole is negative.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The East-West position (degrees; -180 to +180), where 0 is the Prime Meridian (Greenwich), East is positive, West is negative.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude (meters), where 0 is the mean sea level.
	Altitude float64 `protobuf:"fixed64,5,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Accuracy of the location (meters).
	Accuracy float32 `protobuf:"fixed32,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Location) GetAccuracy() float32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

type TimedLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Location *Location              `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *TimedLocation) Reset() {
	*x = TimedLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimedLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimedLocation) ProtoMessage() {}

func (x *TimedLocation) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimedLocation.ProtoReflect.Descriptor instead.
func (*TimedLocation) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_location_proto_rawDescGZIP(), []int{1}
}

func (x *TimedLocation) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *TimedLocation) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_packetbroker_api_v3_location_proto protoreflect.FileDescriptor

var file_packetbroker_api_v3_location_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a,
	0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x7a, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x3b,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_v3_location_proto_rawDescOnce sync.Once
	file_packetbroker_api_v3_location_proto_rawDescData = file_packetbroker_api_v3_location_proto_rawDesc
)

func file_packetbroker_api_v3_location_proto_rawDescGZIP() []byte {
	file_packetbroker_api_v3_location_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_v3_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_v3_location_proto_rawDescData)
	})
	return file_packetbroker_api_v3_location_proto_rawDescData
}

var file_packetbroker_api_v3_location_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_packetbroker_api_v3_location_proto_goTypes = []interface{}{
	(*Location)(nil),              // 0: org.packetbroker.v3.Location
	(*TimedLocation)(nil),         // 1: org.packetbroker.v3.TimedLocation
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_packetbroker_api_v3_location_proto_depIdxs = []int32{
	2, // 0: org.packetbroker.v3.TimedLocation.time:type_name -> google.protobuf.Timestamp
	0, // 1: org.packetbroker.v3.TimedLocation.location:type_name -> org.packetbroker.v3.Location
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_packetbroker_api_v3_location_proto_init() }
func file_packetbroker_api_v3_location_proto_init() {
	if File_packetbroker_api_v3_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_v3_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_packetbroker_api_v3_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimedLocation); i {
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
			RawDescriptor: file_packetbroker_api_v3_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packetbroker_api_v3_location_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_v3_location_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_v3_location_proto_msgTypes,
	}.Build()
	File_packetbroker_api_v3_location_proto = out.File
	file_packetbroker_api_v3_location_proto_rawDesc = nil
	file_packetbroker_api_v3_location_proto_goTypes = nil
	file_packetbroker_api_v3_location_proto_depIdxs = nil
}
