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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: packetbroker/api/v3/contact.proto

package packetbroker

import (
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

type ContactInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Person or organization name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email address.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// URL.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ContactInfo) Reset() {
	*x = ContactInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfo) ProtoMessage() {}

func (x *ContactInfo) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfo.ProtoReflect.Descriptor instead.
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ContactInfoValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ContactInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ContactInfoValue) Reset() {
	*x = ContactInfoValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfoValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfoValue) ProtoMessage() {}

func (x *ContactInfoValue) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfoValue.ProtoReflect.Descriptor instead.
func (*ContactInfoValue) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactInfoValue) GetValue() *ContactInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_packetbroker_api_v3_contact_proto protoreflect.FileDescriptor

var file_packetbroker_api_v3_contact_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x22, 0x49, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x4a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_packetbroker_api_v3_contact_proto_rawDescOnce sync.Once
	file_packetbroker_api_v3_contact_proto_rawDescData = file_packetbroker_api_v3_contact_proto_rawDesc
)

func file_packetbroker_api_v3_contact_proto_rawDescGZIP() []byte {
	file_packetbroker_api_v3_contact_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_v3_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_v3_contact_proto_rawDescData)
	})
	return file_packetbroker_api_v3_contact_proto_rawDescData
}

var file_packetbroker_api_v3_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_packetbroker_api_v3_contact_proto_goTypes = []interface{}{
	(*ContactInfo)(nil),      // 0: org.packetbroker.v3.ContactInfo
	(*ContactInfoValue)(nil), // 1: org.packetbroker.v3.ContactInfoValue
}
var file_packetbroker_api_v3_contact_proto_depIdxs = []int32{
	0, // 0: org.packetbroker.v3.ContactInfoValue.value:type_name -> org.packetbroker.v3.ContactInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_packetbroker_api_v3_contact_proto_init() }
func file_packetbroker_api_v3_contact_proto_init() {
	if File_packetbroker_api_v3_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_v3_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactInfo); i {
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
		file_packetbroker_api_v3_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactInfoValue); i {
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
			RawDescriptor: file_packetbroker_api_v3_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packetbroker_api_v3_contact_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_v3_contact_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_v3_contact_proto_msgTypes,
	}.Build()
	File_packetbroker_api_v3_contact_proto = out.File
	file_packetbroker_api_v3_contact_proto_rawDesc = nil
	file_packetbroker_api_v3_contact_proto_goTypes = nil
	file_packetbroker_api_v3_contact_proto_depIdxs = nil
}
