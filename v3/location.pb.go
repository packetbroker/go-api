// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packetbroker/api/v3/location.proto

package packetbroker

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Location struct {
	// The East-West position (degrees; -180 to +180), where 0 is the Prime Meridian (Greenwich), East is positive, West is negative.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The North–South position (degrees; -90 to +90), where 0 is the equator, North pole is positive, South pole is negative.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude (meters), where 0 is the mean sea level.
	Altitude float32 `protobuf:"fixed32,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Accuracy of the location (meters).
	Accuracy             float32  `protobuf:"fixed32,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()      { *m = Location{} }
func (*Location) ProtoMessage() {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_158c15a02f576fc8, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Location.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return m.Size()
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetAccuracy() float32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

type TimedLocation struct {
	Time                 *types.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Location             *Location        `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TimedLocation) Reset()      { *m = TimedLocation{} }
func (*TimedLocation) ProtoMessage() {}
func (*TimedLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_158c15a02f576fc8, []int{1}
}
func (m *TimedLocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimedLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimedLocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimedLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedLocation.Merge(m, src)
}
func (m *TimedLocation) XXX_Size() int {
	return m.Size()
}
func (m *TimedLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedLocation.DiscardUnknown(m)
}

var xxx_messageInfo_TimedLocation proto.InternalMessageInfo

func (m *TimedLocation) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TimedLocation) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "org.packetbroker.v3.Location")
	golang_proto.RegisterType((*Location)(nil), "org.packetbroker.v3.Location")
	proto.RegisterType((*TimedLocation)(nil), "org.packetbroker.v3.TimedLocation")
	golang_proto.RegisterType((*TimedLocation)(nil), "org.packetbroker.v3.TimedLocation")
}

func init() {
	proto.RegisterFile("packetbroker/api/v3/location.proto", fileDescriptor_158c15a02f576fc8)
}
func init() {
	golang_proto.RegisterFile("packetbroker/api/v3/location.proto", fileDescriptor_158c15a02f576fc8)
}

var fileDescriptor_158c15a02f576fc8 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x21, 0x6c, 0xf2, 0x50,
	0x14, 0x85, 0xdf, 0xe5, 0x27, 0x7f, 0x58, 0xc9, 0x4c, 0x67, 0x08, 0xd9, 0xce, 0x08, 0x66, 0xa8,
	0xd7, 0x04, 0xd4, 0x32, 0x37, 0x3d, 0xd5, 0xa0, 0xe6, 0x4a, 0xe9, 0x9a, 0x86, 0xc2, 0x6b, 0xba,
	0x42, 0xb2, 0x65, 0x02, 0x89, 0x9c, 0x9c, 0x59, 0x32, 0x89, 0x44, 0x22, 0x91, 0x48, 0x24, 0x92,
	0xbe, 0x67, 0x90, 0x48, 0xe4, 0x42, 0x69, 0x3b, 0x96, 0x4c, 0x9e, 0x7b, 0xce, 0xe9, 0xfd, 0x7a,
	0x9f, 0x56, 0x0f, 0x2c, 0xbb, 0xe7, 0x44, 0x9d, 0x50, 0xf4, 0x9c, 0xd0, 0xb0, 0x02, 0xcf, 0x18,
	0xb5, 0x0c, 0x5f, 0xd8, 0x56, 0xe4, 0x89, 0x01, 0x0f, 0x42, 0x11, 0x09, 0xfd, 0x42, 0x84, 0x2e,
	0x3f, 0xcd, 0xf1, 0x51, 0xab, 0x7a, 0xed, 0x0a, 0xe1, 0xfa, 0x8e, 0x91, 0x44, 0x3a, 0xc3, 0x27,
	0x23, 0xf2, 0xfa, 0xce, 0x73, 0x64, 0xf5, 0x83, 0x63, 0xab, 0xfe, 0xa6, 0x95, 0x1e, 0xd2, 0xef,
	0xe8, 0x55, 0xad, 0xe4, 0x5b, 0x91, 0x17, 0x0d, 0xbb, 0x4e, 0x85, 0x6a, 0xd4, 0x20, 0x33, 0xd7,
	0xfa, 0xa5, 0x76, 0xe6, 0x8b, 0x81, 0x7b, 0x34, 0x0b, 0x89, 0xf9, 0x33, 0x38, 0x34, 0x2d, 0x3f,
	0x6d, 0xfe, 0xab, 0x51, 0xa3, 0x60, 0xe6, 0x3a, 0xf1, 0x6c, 0x7b, 0x18, 0x5a, 0xf6, 0x4b, 0xa5,
	0x98, 0x7a, 0xa9, 0xae, 0xbf, 0x6a, 0xe7, 0x6d, 0xaf, 0xef, 0x74, 0x73, 0x04, 0xae, 0x15, 0x0f,
	0x84, 0xc9, 0xfa, 0x72, 0xb3, 0xca, 0x8f, 0xf8, 0x3c, 0xc3, 0xe7, 0xed, 0x0c, 0xdf, 0x4c, 0x72,
	0xfa, 0xad, 0x56, 0xca, 0xce, 0x90, 0x50, 0x95, 0x9b, 0x57, 0xfc, 0x8f, 0x3b, 0xf0, 0x6c, 0x81,
	0x99, 0xc7, 0xef, 0x3f, 0x69, 0x19, 0x83, 0x56, 0x31, 0x68, 0x1d, 0x83, 0x6d, 0x62, 0xb0, 0x6d,
	0x0c, 0xb6, 0x8b, 0xc1, 0xf6, 0x31, 0x68, 0x2c, 0x41, 0x13, 0x09, 0x36, 0x95, 0xa0, 0x99, 0x04,
	0x9b, 0x4b, 0xb0, 0x85, 0x04, 0x5b, 0x4a, 0xd0, 0x4a, 0x82, 0xd6, 0x12, 0x6c, 0x23, 0x41, 0x5b,
	0x09, 0xb6, 0x93, 0xa0, 0xbd, 0x04, 0x1b, 0x2b, 0xb0, 0x89, 0x02, 0xbd, 0x2b, 0xb0, 0x0f, 0x05,
	0xfa, 0x52, 0x60, 0x53, 0x05, 0x36, 0x53, 0xa0, 0xb9, 0x02, 0x2d, 0x14, 0xe8, 0xf1, 0xc6, 0x15,
	0xbf, 0xe9, 0x44, 0xe8, 0xa6, 0x2f, 0x7a, 0x77, 0x3a, 0xef, 0xfc, 0x4f, 0x7e, 0xba, 0xf5, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x6b, 0x54, 0x12, 0x15, 0xfc, 0x01, 0x00, 0x00,
}

func (this *Location) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Location)
	if !ok {
		that2, ok := that.(Location)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Latitude != that1.Latitude {
		return false
	}
	if this.Longitude != that1.Longitude {
		return false
	}
	if this.Altitude != that1.Altitude {
		return false
	}
	if this.Accuracy != that1.Accuracy {
		return false
	}
	return true
}
func (this *TimedLocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TimedLocation)
	if !ok {
		that2, ok := that.(TimedLocation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if !this.Location.Equal(that1.Location) {
		return false
	}
	return true
}
func (m *Location) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Location) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Location) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Accuracy != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Accuracy))))
		i--
		dAtA[i] = 0x25
	}
	if m.Altitude != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Altitude))))
		i--
		dAtA[i] = 0x1d
	}
	if m.Longitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i--
		dAtA[i] = 0x11
	}
	if m.Latitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *TimedLocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedLocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimedLocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Location != nil {
		{
			size, err := m.Location.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Time != nil {
		{
			size, err := m.Time.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedLocation(r randyLocation, easy bool) *Location {
	this := &Location{}
	this.Latitude = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Latitude *= -1
	}
	this.Longitude = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Longitude *= -1
	}
	this.Altitude = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Altitude *= -1
	}
	this.Accuracy = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Accuracy *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedTimedLocation(r randyLocation, easy bool) *TimedLocation {
	this := &TimedLocation{}
	if r.Intn(5) != 0 {
		this.Time = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Location = NewPopulatedLocation(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyLocation interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLocation(r randyLocation) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringLocation(r randyLocation) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneLocation(r)
	}
	return string(tmps)
}
func randUnrecognizedLocation(r randyLocation, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldLocation(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldLocation(dAtA []byte, r randyLocation, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateLocation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateLocation(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Location) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	if m.Altitude != 0 {
		n += 5
	}
	if m.Accuracy != 0 {
		n += 5
	}
	return n
}

func (m *TimedLocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != nil {
		l = m.Time.Size()
		n += 1 + l + sovLocation(uint64(l))
	}
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovLocation(uint64(l))
	}
	return n
}

func sovLocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocation(x uint64) (n int) {
	return sovLocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Location) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Location{`,
		`Latitude:` + fmt.Sprintf("%v", this.Latitude) + `,`,
		`Longitude:` + fmt.Sprintf("%v", this.Longitude) + `,`,
		`Altitude:` + fmt.Sprintf("%v", this.Altitude) + `,`,
		`Accuracy:` + fmt.Sprintf("%v", this.Accuracy) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TimedLocation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TimedLocation{`,
		`Time:` + strings.Replace(fmt.Sprintf("%v", this.Time), "Timestamp", "types.Timestamp", 1) + `,`,
		`Location:` + strings.Replace(this.Location.String(), "Location", "Location", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLocation(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Location) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocation
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
			return fmt.Errorf("proto: Location: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Location: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Altitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Altitude = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accuracy", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Accuracy = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipLocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocation
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
func (m *TimedLocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocation
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
			return fmt.Errorf("proto: TimedLocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedLocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocation
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
				return ErrInvalidLengthLocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = &types.Timestamp{}
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocation
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
				return ErrInvalidLengthLocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Location == nil {
				m.Location = &Location{}
			}
			if err := m.Location.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocation
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
func skipLocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocation
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
					return 0, ErrIntOverflowLocation
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
					return 0, ErrIntOverflowLocation
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
				return 0, ErrInvalidLengthLocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocation = fmt.Errorf("proto: unexpected end of group")
)
