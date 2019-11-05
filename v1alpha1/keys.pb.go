// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packetbroker/api/v1alpha1/keys.proto

package packetbroker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Pointer to a Key on a Key Exchange.
type KeyPointer struct {
	// Address of the Key Exchange.
	KeyExchange string `protobuf:"bytes,1,opt,name=key_exchange,json=keyExchange,proto3" json:"key_exchange,omitempty"`
	// Label of the Key.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *KeyPointer) Reset()         { *m = KeyPointer{} }
func (m *KeyPointer) String() string { return proto.CompactTextString(m) }
func (*KeyPointer) ProtoMessage()    {}
func (*KeyPointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf5357c5d765333, []int{0}
}
func (m *KeyPointer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyPointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyPointer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyPointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPointer.Merge(m, src)
}
func (m *KeyPointer) XXX_Size() int {
	return m.Size()
}
func (m *KeyPointer) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPointer.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPointer proto.InternalMessageInfo

func (m *KeyPointer) GetKeyExchange() string {
	if m != nil {
		return m.KeyExchange
	}
	return ""
}

func (m *KeyPointer) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// Encrypted data.
type EncryptedData struct {
	// Encrypted data encryption keys with which the value is encrypted using AES-256-GCM.
	// The map key refers to the key encryption key.
	DeksEncrypted map[string][]byte `protobuf:"bytes,1,rep,name=deks_encrypted,json=deksEncrypted,proto3" json:"deks_encrypted,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Encrypted data.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *EncryptedData) Reset()         { *m = EncryptedData{} }
func (m *EncryptedData) String() string { return proto.CompactTextString(m) }
func (*EncryptedData) ProtoMessage()    {}
func (*EncryptedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bf5357c5d765333, []int{1}
}
func (m *EncryptedData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncryptedData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncryptedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedData.Merge(m, src)
}
func (m *EncryptedData) XXX_Size() int {
	return m.Size()
}
func (m *EncryptedData) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedData.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedData proto.InternalMessageInfo

func (m *EncryptedData) GetDeksEncrypted() map[string][]byte {
	if m != nil {
		return m.DeksEncrypted
	}
	return nil
}

func (m *EncryptedData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyPointer)(nil), "org.packetbroker.v1alpha1.KeyPointer")
	proto.RegisterType((*EncryptedData)(nil), "org.packetbroker.v1alpha1.EncryptedData")
	proto.RegisterMapType((map[string][]byte)(nil), "org.packetbroker.v1alpha1.EncryptedData.DeksEncryptedEntry")
}

func init() {
	proto.RegisterFile("packetbroker/api/v1alpha1/keys.proto", fileDescriptor_1bf5357c5d765333)
}

var fileDescriptor_1bf5357c5d765333 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0x49, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x4c,
	0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0xcf, 0x4e, 0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0xcc, 0x2f, 0x4a, 0xd7, 0x43, 0x56, 0xa9, 0x07, 0x53, 0xa5, 0xe4, 0xca, 0xc5, 0xe5,
	0x9d, 0x5a, 0x19, 0x90, 0x9f, 0x99, 0x57, 0x92, 0x5a, 0x24, 0xa4, 0xc8, 0xc5, 0x93, 0x9d, 0x5a,
	0x19, 0x9f, 0x5a, 0x91, 0x9c, 0x91, 0x98, 0x97, 0x9e, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0xc4, 0x9d, 0x9d, 0x5a, 0xe9, 0x0a, 0x15, 0x12, 0x12, 0xe1, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd,
	0x91, 0x60, 0x02, 0xcb, 0x41, 0x38, 0x4a, 0xa7, 0x19, 0xb9, 0x78, 0x5d, 0xf3, 0x92, 0x8b, 0x2a,
	0x0b, 0x4a, 0x52, 0x53, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x92, 0xb8, 0xf8, 0x52, 0x52, 0xb3, 0x8b,
	0xe3, 0x53, 0x61, 0xa2, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xd6, 0x7a, 0x38, 0x1d, 0xa3,
	0x87, 0x62, 0x82, 0x9e, 0x4b, 0x6a, 0x76, 0x31, 0x5c, 0xc4, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x88,
	0x37, 0x05, 0x59, 0x0c, 0xe4, 0x96, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0xb0, 0x5b, 0x78, 0x82, 0x20,
	0x1c, 0x29, 0x07, 0x2e, 0x21, 0x4c, 0xad, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x1f,
	0x81, 0x98, 0xd8, 0x75, 0x5b, 0x31, 0x59, 0x30, 0x3a, 0xb9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a, 0x3e, 0xaa, 0xdb, 0xf3, 0x8b, 0xd2, 0x51, 0x82, 0xdd,
	0x1a, 0x59, 0x36, 0x89, 0x0d, 0x1c, 0xfe, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x14,
	0x94, 0xb2, 0xa7, 0x01, 0x00, 0x00,
}

func (m *KeyPointer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyPointer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyPointer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KeyExchange) > 0 {
		i -= len(m.KeyExchange)
		copy(dAtA[i:], m.KeyExchange)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.KeyExchange)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EncryptedData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintKeys(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeksEncrypted) > 0 {
		for k := range m.DeksEncrypted {
			v := m.DeksEncrypted[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintKeys(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintKeys(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintKeys(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyPointer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyExchange)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *EncryptedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DeksEncrypted) > 0 {
		for k, v := range m.DeksEncrypted {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovKeys(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovKeys(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovKeys(uint64(mapEntrySize))
		}
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyPointer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyPointer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyPointer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyExchange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyExchange = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EncryptedData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EncryptedData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptedData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeksEncrypted", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeksEncrypted == nil {
				m.DeksEncrypted = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowKeys
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowKeys
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthKeys
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthKeys
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowKeys
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthKeys
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthKeys
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipKeys(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthKeys
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DeksEncrypted[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)
