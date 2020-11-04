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
// 	protoc        v3.13.0
// source: packetbroker/api/v3/keys.proto

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

// Pointer to a Key on a Key Exchange.
type KeyPointer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of the Key Exchange.
	KeyExchange string `protobuf:"bytes,1,opt,name=key_exchange,json=keyExchange,proto3" json:"key_exchange,omitempty"`
	// Label of the Key.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *KeyPointer) Reset() {
	*x = KeyPointer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyPointer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyPointer) ProtoMessage() {}

func (x *KeyPointer) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyPointer.ProtoReflect.Descriptor instead.
func (*KeyPointer) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_keys_proto_rawDescGZIP(), []int{0}
}

func (x *KeyPointer) GetKeyExchange() string {
	if x != nil {
		return x.KeyExchange
	}
	return ""
}

func (x *KeyPointer) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Encrypted data.
type EncryptedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted data encryption keys with which the value is encrypted using AES-256-GCM.
	// The map key refers to the key encryption key.
	DeksEncrypted map[string][]byte `protobuf:"bytes,1,rep,name=deks_encrypted,json=deksEncrypted,proto3" json:"deks_encrypted,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Encrypted data.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EncryptedData) Reset() {
	*x = EncryptedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packetbroker_api_v3_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedData) ProtoMessage() {}

func (x *EncryptedData) ProtoReflect() protoreflect.Message {
	mi := &file_packetbroker_api_v3_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedData.ProtoReflect.Descriptor instead.
func (*EncryptedData) Descriptor() ([]byte, []int) {
	return file_packetbroker_api_v3_keys_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptedData) GetDeksEncrypted() map[string][]byte {
	if x != nil {
		return x.DeksEncrypted
	}
	return nil
}

func (x *EncryptedData) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_packetbroker_api_v3_keys_proto protoreflect.FileDescriptor

var file_packetbroker_api_v3_keys_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x22, 0x45, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xc5, 0x01, 0x0a,
	0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5c,
	0x0a, 0x0e, 0x64, 0x65, 0x6b, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6b, 0x73, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x64,
	0x65, 0x6b, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x40, 0x0a, 0x12, 0x44, 0x65, 0x6b, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x33, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packetbroker_api_v3_keys_proto_rawDescOnce sync.Once
	file_packetbroker_api_v3_keys_proto_rawDescData = file_packetbroker_api_v3_keys_proto_rawDesc
)

func file_packetbroker_api_v3_keys_proto_rawDescGZIP() []byte {
	file_packetbroker_api_v3_keys_proto_rawDescOnce.Do(func() {
		file_packetbroker_api_v3_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_packetbroker_api_v3_keys_proto_rawDescData)
	})
	return file_packetbroker_api_v3_keys_proto_rawDescData
}

var file_packetbroker_api_v3_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packetbroker_api_v3_keys_proto_goTypes = []interface{}{
	(*KeyPointer)(nil),    // 0: org.packetbroker.v3.KeyPointer
	(*EncryptedData)(nil), // 1: org.packetbroker.v3.EncryptedData
	nil,                   // 2: org.packetbroker.v3.EncryptedData.DeksEncryptedEntry
}
var file_packetbroker_api_v3_keys_proto_depIdxs = []int32{
	2, // 0: org.packetbroker.v3.EncryptedData.deks_encrypted:type_name -> org.packetbroker.v3.EncryptedData.DeksEncryptedEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_packetbroker_api_v3_keys_proto_init() }
func file_packetbroker_api_v3_keys_proto_init() {
	if File_packetbroker_api_v3_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packetbroker_api_v3_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyPointer); i {
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
		file_packetbroker_api_v3_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedData); i {
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
			RawDescriptor: file_packetbroker_api_v3_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packetbroker_api_v3_keys_proto_goTypes,
		DependencyIndexes: file_packetbroker_api_v3_keys_proto_depIdxs,
		MessageInfos:      file_packetbroker_api_v3_keys_proto_msgTypes,
	}.Build()
	File_packetbroker_api_v3_keys_proto = out.File
	file_packetbroker_api_v3_keys_proto_rawDesc = nil
	file_packetbroker_api_v3_keys_proto_goTypes = nil
	file_packetbroker_api_v3_keys_proto_depIdxs = nil
}
