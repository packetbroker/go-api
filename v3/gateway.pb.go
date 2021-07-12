// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packetbroker/api/v3/gateway.proto

package packetbroker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
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

type GatewayVisibility struct {
	// LoRa Alliance NetID of the Forwarder Member.
	ForwarderNetId uint32 `protobuf:"varint,10,opt,name=forwarder_net_id,json=forwarderNetId,proto3" json:"forwarder_net_id,omitempty"`
	// Tenant ID managed by the Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,11,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
	// LoRa Alliance NetID of the Home Network Member.
	HomeNetworkNetId uint32 `protobuf:"varint,12,opt,name=home_network_net_id,json=homeNetworkNetId,proto3" json:"home_network_net_id,omitempty"`
	// Tenant ID managed by the Home Network Member.
	HomeNetworkTenantId string `protobuf:"bytes,13,opt,name=home_network_tenant_id,json=homeNetworkTenantId,proto3" json:"home_network_tenant_id,omitempty"`
	// Timestamp when the policy got last updated.
	UpdatedAt *types.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Show location.
	Location bool `protobuf:"varint,1,opt,name=location,proto3" json:"location,omitempty"`
	// Show antenna placement (indoor/outdoor).
	AntennaPlacement bool `protobuf:"varint,2,opt,name=antenna_placement,json=antennaPlacement,proto3" json:"antenna_placement,omitempty"`
	// Show antenna count.
	AntennaCount bool `protobuf:"varint,3,opt,name=antenna_count,json=antennaCount,proto3" json:"antenna_count,omitempty"`
	// Show whether the gateway produces fine timestamps.
	FineTimestamps bool `protobuf:"varint,4,opt,name=fine_timestamps,json=fineTimestamps,proto3" json:"fine_timestamps,omitempty"`
	// Show contact information.
	ContactInfo bool `protobuf:"varint,5,opt,name=contact_info,json=contactInfo,proto3" json:"contact_info,omitempty"`
	// Show status (online/offline).
	Status bool `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	// Show frequency plan.
	FrequencyPlan bool `protobuf:"varint,8,opt,name=frequency_plan,json=frequencyPlan,proto3" json:"frequency_plan,omitempty"`
	// Show receive and transmission packet rates.
	PacketRates          bool     `protobuf:"varint,9,opt,name=packet_rates,json=packetRates,proto3" json:"packet_rates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayVisibility) Reset()         { *m = GatewayVisibility{} }
func (m *GatewayVisibility) String() string { return proto.CompactTextString(m) }
func (*GatewayVisibility) ProtoMessage()    {}
func (*GatewayVisibility) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{0}
}
func (m *GatewayVisibility) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayVisibility.Unmarshal(m, b)
}
func (m *GatewayVisibility) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayVisibility.Marshal(b, m, deterministic)
}
func (m *GatewayVisibility) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayVisibility.Merge(m, src)
}
func (m *GatewayVisibility) XXX_Size() int {
	return xxx_messageInfo_GatewayVisibility.Size(m)
}
func (m *GatewayVisibility) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayVisibility.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayVisibility proto.InternalMessageInfo

func (m *GatewayVisibility) GetForwarderNetId() uint32 {
	if m != nil {
		return m.ForwarderNetId
	}
	return 0
}

func (m *GatewayVisibility) GetForwarderTenantId() string {
	if m != nil {
		return m.ForwarderTenantId
	}
	return ""
}

func (m *GatewayVisibility) GetHomeNetworkNetId() uint32 {
	if m != nil {
		return m.HomeNetworkNetId
	}
	return 0
}

func (m *GatewayVisibility) GetHomeNetworkTenantId() string {
	if m != nil {
		return m.HomeNetworkTenantId
	}
	return ""
}

func (m *GatewayVisibility) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *GatewayVisibility) GetLocation() bool {
	if m != nil {
		return m.Location
	}
	return false
}

func (m *GatewayVisibility) GetAntennaPlacement() bool {
	if m != nil {
		return m.AntennaPlacement
	}
	return false
}

func (m *GatewayVisibility) GetAntennaCount() bool {
	if m != nil {
		return m.AntennaCount
	}
	return false
}

func (m *GatewayVisibility) GetFineTimestamps() bool {
	if m != nil {
		return m.FineTimestamps
	}
	return false
}

func (m *GatewayVisibility) GetContactInfo() bool {
	if m != nil {
		return m.ContactInfo
	}
	return false
}

func (m *GatewayVisibility) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GatewayVisibility) GetFrequencyPlan() bool {
	if m != nil {
		return m.FrequencyPlan
	}
	return false
}

func (m *GatewayVisibility) GetPacketRates() bool {
	if m != nil {
		return m.PacketRates
	}
	return false
}

type GatewayFrequencyPlan struct {
	// Region of the frequency plan.
	Region Region `protobuf:"varint,1,opt,name=region,proto3,enum=org.packetbroker.v3.Region" json:"region,omitempty"`
	// LoRa channels demodulating multiple spreading factors per channel.
	LoraMultiSfChannels []*GatewayFrequencyPlan_LoRaMultiSFChannel `protobuf:"bytes,3,rep,name=lora_multi_sf_channels,json=loraMultiSfChannels,proto3" json:"lora_multi_sf_channels,omitempty"`
	// LoRa channels for a single spreading factor and bandwidth.
	// This is typically used for LoRa 2.4 GHz gateways.
	LoraSingleSfChannels []*GatewayFrequencyPlan_LoRaSingleSFChannel `protobuf:"bytes,4,rep,name=lora_single_sf_channels,json=loraSingleSfChannels,proto3" json:"lora_single_sf_channels,omitempty"`
	// FSK channel.
	FskChannel           *GatewayFrequencyPlan_FSKChannel `protobuf:"bytes,5,opt,name=fsk_channel,json=fskChannel,proto3" json:"fsk_channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GatewayFrequencyPlan) Reset()         { *m = GatewayFrequencyPlan{} }
func (m *GatewayFrequencyPlan) String() string { return proto.CompactTextString(m) }
func (*GatewayFrequencyPlan) ProtoMessage()    {}
func (*GatewayFrequencyPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{1}
}
func (m *GatewayFrequencyPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayFrequencyPlan.Unmarshal(m, b)
}
func (m *GatewayFrequencyPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayFrequencyPlan.Marshal(b, m, deterministic)
}
func (m *GatewayFrequencyPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayFrequencyPlan.Merge(m, src)
}
func (m *GatewayFrequencyPlan) XXX_Size() int {
	return xxx_messageInfo_GatewayFrequencyPlan.Size(m)
}
func (m *GatewayFrequencyPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayFrequencyPlan.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayFrequencyPlan proto.InternalMessageInfo

func (m *GatewayFrequencyPlan) GetRegion() Region {
	if m != nil {
		return m.Region
	}
	return Region_UNKNOWN_REGION
}

func (m *GatewayFrequencyPlan) GetLoraMultiSfChannels() []*GatewayFrequencyPlan_LoRaMultiSFChannel {
	if m != nil {
		return m.LoraMultiSfChannels
	}
	return nil
}

func (m *GatewayFrequencyPlan) GetLoraSingleSfChannels() []*GatewayFrequencyPlan_LoRaSingleSFChannel {
	if m != nil {
		return m.LoraSingleSfChannels
	}
	return nil
}

func (m *GatewayFrequencyPlan) GetFskChannel() *GatewayFrequencyPlan_FSKChannel {
	if m != nil {
		return m.FskChannel
	}
	return nil
}

type GatewayFrequencyPlan_LoRaMultiSFChannel struct {
	// Frequency (Hz).
	Frequency            uint64   `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) Reset() {
	*m = GatewayFrequencyPlan_LoRaMultiSFChannel{}
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) String() string { return proto.CompactTextString(m) }
func (*GatewayFrequencyPlan_LoRaMultiSFChannel) ProtoMessage()    {}
func (*GatewayFrequencyPlan_LoRaMultiSFChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{1, 0}
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel.Unmarshal(m, b)
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel.Marshal(b, m, deterministic)
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel.Merge(m, src)
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) XXX_Size() int {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel.Size(m)
}
func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayFrequencyPlan_LoRaMultiSFChannel proto.InternalMessageInfo

func (m *GatewayFrequencyPlan_LoRaMultiSFChannel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type GatewayFrequencyPlan_LoRaSingleSFChannel struct {
	// Frequency (Hz).
	Frequency uint64 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// LoRa spreading factor.
	SpreadingFactor uint32 `protobuf:"varint,2,opt,name=spreading_factor,json=spreadingFactor,proto3" json:"spreading_factor,omitempty"`
	// LoRa bandwidth (Hz).
	Bandwidth            uint32   `protobuf:"varint,3,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) Reset() {
	*m = GatewayFrequencyPlan_LoRaSingleSFChannel{}
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) String() string { return proto.CompactTextString(m) }
func (*GatewayFrequencyPlan_LoRaSingleSFChannel) ProtoMessage()    {}
func (*GatewayFrequencyPlan_LoRaSingleSFChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{1, 1}
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel.Unmarshal(m, b)
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel.Marshal(b, m, deterministic)
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel.Merge(m, src)
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) XXX_Size() int {
	return xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel.Size(m)
}
func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayFrequencyPlan_LoRaSingleSFChannel proto.InternalMessageInfo

func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) GetSpreadingFactor() uint32 {
	if m != nil {
		return m.SpreadingFactor
	}
	return 0
}

func (m *GatewayFrequencyPlan_LoRaSingleSFChannel) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

type GatewayFrequencyPlan_FSKChannel struct {
	// Frequency (Hz).
	Frequency            uint64   `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayFrequencyPlan_FSKChannel) Reset()         { *m = GatewayFrequencyPlan_FSKChannel{} }
func (m *GatewayFrequencyPlan_FSKChannel) String() string { return proto.CompactTextString(m) }
func (*GatewayFrequencyPlan_FSKChannel) ProtoMessage()    {}
func (*GatewayFrequencyPlan_FSKChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{1, 2}
}
func (m *GatewayFrequencyPlan_FSKChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayFrequencyPlan_FSKChannel.Unmarshal(m, b)
}
func (m *GatewayFrequencyPlan_FSKChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayFrequencyPlan_FSKChannel.Marshal(b, m, deterministic)
}
func (m *GatewayFrequencyPlan_FSKChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayFrequencyPlan_FSKChannel.Merge(m, src)
}
func (m *GatewayFrequencyPlan_FSKChannel) XXX_Size() int {
	return xxx_messageInfo_GatewayFrequencyPlan_FSKChannel.Size(m)
}
func (m *GatewayFrequencyPlan_FSKChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayFrequencyPlan_FSKChannel.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayFrequencyPlan_FSKChannel proto.InternalMessageInfo

func (m *GatewayFrequencyPlan_FSKChannel) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type GatewayIdentifier struct {
	// Extended unique identifier of the gateway.
	Eui *types.UInt64Value `protobuf:"bytes,1,opt,name=eui,proto3" json:"eui,omitempty"`
	// Types that are valid to be assigned to Id:
	//	*GatewayIdentifier_Plain
	//	*GatewayIdentifier_Hash
	Id                   isGatewayIdentifier_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GatewayIdentifier) Reset()         { *m = GatewayIdentifier{} }
func (m *GatewayIdentifier) String() string { return proto.CompactTextString(m) }
func (*GatewayIdentifier) ProtoMessage()    {}
func (*GatewayIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{2}
}
func (m *GatewayIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayIdentifier.Unmarshal(m, b)
}
func (m *GatewayIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayIdentifier.Marshal(b, m, deterministic)
}
func (m *GatewayIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayIdentifier.Merge(m, src)
}
func (m *GatewayIdentifier) XXX_Size() int {
	return xxx_messageInfo_GatewayIdentifier.Size(m)
}
func (m *GatewayIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayIdentifier proto.InternalMessageInfo

type isGatewayIdentifier_Id interface {
	isGatewayIdentifier_Id()
}

type GatewayIdentifier_Plain struct {
	Plain string `protobuf:"bytes,2,opt,name=plain,proto3,oneof" json:"plain,omitempty"`
}
type GatewayIdentifier_Hash struct {
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3,oneof" json:"hash,omitempty"`
}

func (*GatewayIdentifier_Plain) isGatewayIdentifier_Id() {}
func (*GatewayIdentifier_Hash) isGatewayIdentifier_Id()  {}

func (m *GatewayIdentifier) GetId() isGatewayIdentifier_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GatewayIdentifier) GetEui() *types.UInt64Value {
	if m != nil {
		return m.Eui
	}
	return nil
}

func (m *GatewayIdentifier) GetPlain() string {
	if x, ok := m.GetId().(*GatewayIdentifier_Plain); ok {
		return x.Plain
	}
	return ""
}

func (m *GatewayIdentifier) GetHash() []byte {
	if x, ok := m.GetId().(*GatewayIdentifier_Hash); ok {
		return x.Hash
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GatewayIdentifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GatewayIdentifier_Plain)(nil),
		(*GatewayIdentifier_Hash)(nil),
	}
}

type GatewayLocation struct {
	// Types that are valid to be assigned to Value:
	//	*GatewayLocation_Terrestrial_
	//	*GatewayLocation_Satellite_
	Value                isGatewayLocation_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GatewayLocation) Reset()         { *m = GatewayLocation{} }
func (m *GatewayLocation) String() string { return proto.CompactTextString(m) }
func (*GatewayLocation) ProtoMessage()    {}
func (*GatewayLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{3}
}
func (m *GatewayLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayLocation.Unmarshal(m, b)
}
func (m *GatewayLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayLocation.Marshal(b, m, deterministic)
}
func (m *GatewayLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayLocation.Merge(m, src)
}
func (m *GatewayLocation) XXX_Size() int {
	return xxx_messageInfo_GatewayLocation.Size(m)
}
func (m *GatewayLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayLocation proto.InternalMessageInfo

type isGatewayLocation_Value interface {
	isGatewayLocation_Value()
}

type GatewayLocation_Terrestrial_ struct {
	Terrestrial *GatewayLocation_Terrestrial `protobuf:"bytes,1,opt,name=terrestrial,proto3,oneof" json:"terrestrial,omitempty"`
}
type GatewayLocation_Satellite_ struct {
	Satellite *GatewayLocation_Satellite `protobuf:"bytes,2,opt,name=satellite,proto3,oneof" json:"satellite,omitempty"`
}

func (*GatewayLocation_Terrestrial_) isGatewayLocation_Value() {}
func (*GatewayLocation_Satellite_) isGatewayLocation_Value()   {}

func (m *GatewayLocation) GetValue() isGatewayLocation_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GatewayLocation) GetTerrestrial() *GatewayLocation_Terrestrial {
	if x, ok := m.GetValue().(*GatewayLocation_Terrestrial_); ok {
		return x.Terrestrial
	}
	return nil
}

func (m *GatewayLocation) GetSatellite() *GatewayLocation_Satellite {
	if x, ok := m.GetValue().(*GatewayLocation_Satellite_); ok {
		return x.Satellite
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GatewayLocation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GatewayLocation_Terrestrial_)(nil),
		(*GatewayLocation_Satellite_)(nil),
	}
}

type GatewayLocation_Terrestrial struct {
	// Location of the (first) antenna.
	// Home Networks receive this value if the Gateway Visibility has location set.
	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// Antenna placement (indoor/outdoor).
	// Home Networks receive this value if the Gateway Visibility has antenna_placement set.
	AntennaPlacement TerrestrialAntennaPlacement `protobuf:"varint,2,opt,name=antenna_placement,json=antennaPlacement,proto3,enum=org.packetbroker.v3.TerrestrialAntennaPlacement" json:"antenna_placement,omitempty"`
	// Number of antennas.
	// Home Networks receive this value if the Gateway Visibility has antenna_count set.
	AntennaCount *types.UInt32Value `protobuf:"bytes,3,opt,name=antenna_count,json=antennaCount,proto3" json:"antenna_count,omitempty"`
	// Indicates whether the gateway produces fine timestamps.
	// Home Networks receive this value if the Gateway Visibility has fine_timestamps set.
	FineTimestamps       *types.BoolValue `protobuf:"bytes,4,opt,name=fine_timestamps,json=fineTimestamps,proto3" json:"fine_timestamps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GatewayLocation_Terrestrial) Reset()         { *m = GatewayLocation_Terrestrial{} }
func (m *GatewayLocation_Terrestrial) String() string { return proto.CompactTextString(m) }
func (*GatewayLocation_Terrestrial) ProtoMessage()    {}
func (*GatewayLocation_Terrestrial) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{3, 0}
}
func (m *GatewayLocation_Terrestrial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayLocation_Terrestrial.Unmarshal(m, b)
}
func (m *GatewayLocation_Terrestrial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayLocation_Terrestrial.Marshal(b, m, deterministic)
}
func (m *GatewayLocation_Terrestrial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayLocation_Terrestrial.Merge(m, src)
}
func (m *GatewayLocation_Terrestrial) XXX_Size() int {
	return xxx_messageInfo_GatewayLocation_Terrestrial.Size(m)
}
func (m *GatewayLocation_Terrestrial) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayLocation_Terrestrial.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayLocation_Terrestrial proto.InternalMessageInfo

func (m *GatewayLocation_Terrestrial) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *GatewayLocation_Terrestrial) GetAntennaPlacement() TerrestrialAntennaPlacement {
	if m != nil {
		return m.AntennaPlacement
	}
	return TerrestrialAntennaPlacement_UNKNOWN_PLACEMENT
}

func (m *GatewayLocation_Terrestrial) GetAntennaCount() *types.UInt32Value {
	if m != nil {
		return m.AntennaCount
	}
	return nil
}

func (m *GatewayLocation_Terrestrial) GetFineTimestamps() *types.BoolValue {
	if m != nil {
		return m.FineTimestamps
	}
	return nil
}

type GatewayLocation_Satellite struct {
	// Satellite location.
	// Home Networks receive this value if the Gateway Visibility has location set.
	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// Field of view (meters).
	// Home Networks receive this value if the Gateway Visibility has location set.
	FieldOfView float64 `protobuf:"fixed64,2,opt,name=field_of_view,json=fieldOfView,proto3" json:"field_of_view,omitempty"`
	// Trajectory, typically containing the recent past and projected path.
	// Home Networks receive this value if the Gateway Visibility has location set.
	Trajectory           []*TimedLocation `protobuf:"bytes,3,rep,name=trajectory,proto3" json:"trajectory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GatewayLocation_Satellite) Reset()         { *m = GatewayLocation_Satellite{} }
func (m *GatewayLocation_Satellite) String() string { return proto.CompactTextString(m) }
func (*GatewayLocation_Satellite) ProtoMessage()    {}
func (*GatewayLocation_Satellite) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{3, 1}
}
func (m *GatewayLocation_Satellite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayLocation_Satellite.Unmarshal(m, b)
}
func (m *GatewayLocation_Satellite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayLocation_Satellite.Marshal(b, m, deterministic)
}
func (m *GatewayLocation_Satellite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayLocation_Satellite.Merge(m, src)
}
func (m *GatewayLocation_Satellite) XXX_Size() int {
	return xxx_messageInfo_GatewayLocation_Satellite.Size(m)
}
func (m *GatewayLocation_Satellite) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayLocation_Satellite.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayLocation_Satellite proto.InternalMessageInfo

func (m *GatewayLocation_Satellite) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *GatewayLocation_Satellite) GetFieldOfView() float64 {
	if m != nil {
		return m.FieldOfView
	}
	return 0
}

func (m *GatewayLocation_Satellite) GetTrajectory() []*TimedLocation {
	if m != nil {
		return m.Trajectory
	}
	return nil
}

type GatewayLocationValue struct {
	Location             *GatewayLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GatewayLocationValue) Reset()         { *m = GatewayLocationValue{} }
func (m *GatewayLocationValue) String() string { return proto.CompactTextString(m) }
func (*GatewayLocationValue) ProtoMessage()    {}
func (*GatewayLocationValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{4}
}
func (m *GatewayLocationValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayLocationValue.Unmarshal(m, b)
}
func (m *GatewayLocationValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayLocationValue.Marshal(b, m, deterministic)
}
func (m *GatewayLocationValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayLocationValue.Merge(m, src)
}
func (m *GatewayLocationValue) XXX_Size() int {
	return xxx_messageInfo_GatewayLocationValue.Size(m)
}
func (m *GatewayLocationValue) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayLocationValue.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayLocationValue proto.InternalMessageInfo

func (m *GatewayLocationValue) GetLocation() *GatewayLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

type Gateway struct {
	// LoRa Alliance NetID of the Forwarder Member.
	NetId uint32 `protobuf:"varint,1,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	// Tenant ID managed by the Forwarder Member.
	TenantId string `protobuf:"bytes,14,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Forwarder cluster of the Forwarder Member (optional).
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Identifier of the gateway.
	GatewayId *GatewayIdentifier `protobuf:"bytes,15,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// Timestamp when the gateway was last updated.
	UpdatedAt *types.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Gateway location.
	Location *GatewayLocation `protobuf:"bytes,20,opt,name=location,proto3" json:"location,omitempty"`
	// Administrative contact.
	// Home Networks receive this value if the Gateway Visibility has contact_info set.
	AdministrativeContact *ContactInfo `protobuf:"bytes,16,opt,name=administrative_contact,json=administrativeContact,proto3" json:"administrative_contact,omitempty"`
	// Technical contact.
	// Home Networks receive this value if the Gateway Visibility has contact_info set.
	TechnicalContact *ContactInfo `protobuf:"bytes,17,opt,name=technical_contact,json=technicalContact,proto3" json:"technical_contact,omitempty"`
	// Indicates whether the gateway is online.
	// Home Networks receive this value if the Gateway Visibility has status set.
	Online *types.BoolValue `protobuf:"bytes,10,opt,name=online,proto3" json:"online,omitempty"`
	// Frequency plan of the gateway.
	// Home Networks receive this value if the Gateway Visibility has frequency_plan set.
	FrequencyPlan *GatewayFrequencyPlan `protobuf:"bytes,13,opt,name=frequency_plan,json=frequencyPlan,proto3" json:"frequency_plan,omitempty"`
	// Received packets rate (number of packets per hour).
	// Home Networks receive this value if the Gateway Visibility has packet_rates set.
	RxRate *types.FloatValue `protobuf:"bytes,18,opt,name=rx_rate,json=rxRate,proto3" json:"rx_rate,omitempty"`
	// Transmitted packets rate (number of packets per hour).
	// Home Networks receive this value if the Gateway Visibility has packet_rates set.
	TxRate               *types.FloatValue `protobuf:"bytes,19,opt,name=tx_rate,json=txRate,proto3" json:"tx_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f115301d3f8940, []int{5}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetNetId() uint32 {
	if m != nil {
		return m.NetId
	}
	return 0
}

func (m *Gateway) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *Gateway) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Gateway) GetGatewayId() *GatewayIdentifier {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *Gateway) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Gateway) GetLocation() *GatewayLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Gateway) GetAdministrativeContact() *ContactInfo {
	if m != nil {
		return m.AdministrativeContact
	}
	return nil
}

func (m *Gateway) GetTechnicalContact() *ContactInfo {
	if m != nil {
		return m.TechnicalContact
	}
	return nil
}

func (m *Gateway) GetOnline() *types.BoolValue {
	if m != nil {
		return m.Online
	}
	return nil
}

func (m *Gateway) GetFrequencyPlan() *GatewayFrequencyPlan {
	if m != nil {
		return m.FrequencyPlan
	}
	return nil
}

func (m *Gateway) GetRxRate() *types.FloatValue {
	if m != nil {
		return m.RxRate
	}
	return nil
}

func (m *Gateway) GetTxRate() *types.FloatValue {
	if m != nil {
		return m.TxRate
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayVisibility)(nil), "org.packetbroker.v3.GatewayVisibility")
	proto.RegisterType((*GatewayFrequencyPlan)(nil), "org.packetbroker.v3.GatewayFrequencyPlan")
	proto.RegisterType((*GatewayFrequencyPlan_LoRaMultiSFChannel)(nil), "org.packetbroker.v3.GatewayFrequencyPlan.LoRaMultiSFChannel")
	proto.RegisterType((*GatewayFrequencyPlan_LoRaSingleSFChannel)(nil), "org.packetbroker.v3.GatewayFrequencyPlan.LoRaSingleSFChannel")
	proto.RegisterType((*GatewayFrequencyPlan_FSKChannel)(nil), "org.packetbroker.v3.GatewayFrequencyPlan.FSKChannel")
	proto.RegisterType((*GatewayIdentifier)(nil), "org.packetbroker.v3.GatewayIdentifier")
	proto.RegisterType((*GatewayLocation)(nil), "org.packetbroker.v3.GatewayLocation")
	proto.RegisterType((*GatewayLocation_Terrestrial)(nil), "org.packetbroker.v3.GatewayLocation.Terrestrial")
	proto.RegisterType((*GatewayLocation_Satellite)(nil), "org.packetbroker.v3.GatewayLocation.Satellite")
	proto.RegisterType((*GatewayLocationValue)(nil), "org.packetbroker.v3.GatewayLocationValue")
	proto.RegisterType((*Gateway)(nil), "org.packetbroker.v3.Gateway")
}

func init() { proto.RegisterFile("packetbroker/api/v3/gateway.proto", fileDescriptor_56f115301d3f8940) }

var fileDescriptor_56f115301d3f8940 = []byte{
	// 1185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xed, 0x6e, 0xdb, 0x36,
	0x14, 0x8d, 0x13, 0xd9, 0x91, 0xaf, 0x62, 0x47, 0xa1, 0xd3, 0x4c, 0x70, 0xdb, 0x35, 0xf5, 0x3e,
	0x9a, 0x6e, 0x98, 0x32, 0x38, 0xc5, 0x80, 0x62, 0x1b, 0xb0, 0x26, 0x58, 0x96, 0x78, 0x6d, 0x57,
	0x28, 0x69, 0x37, 0x0c, 0x18, 0x04, 0x46, 0xa2, 0x1c, 0x2e, 0x32, 0xe9, 0x52, 0x74, 0xdc, 0x00,
	0x7b, 0x9b, 0xfd, 0xef, 0xde, 0x6b, 0x7f, 0xf6, 0x0a, 0x83, 0x28, 0x4a, 0x72, 0x12, 0x35, 0x70,
	0xb7, 0x7f, 0xd2, 0xb9, 0xe7, 0x9e, 0x73, 0xf9, 0x75, 0x49, 0xb8, 0x3f, 0xc6, 0xc1, 0x19, 0x91,
	0x27, 0x82, 0x9f, 0x11, 0xb1, 0x8d, 0xc7, 0x74, 0xfb, 0x7c, 0x67, 0x7b, 0x88, 0x25, 0x99, 0xe2,
	0x0b, 0x77, 0x2c, 0xb8, 0xe4, 0xa8, 0xc3, 0xc5, 0xd0, 0x9d, 0xa5, 0xb9, 0xe7, 0x3b, 0xdd, 0x7b,
	0x43, 0xce, 0x87, 0x31, 0xd9, 0x56, 0x94, 0x93, 0x49, 0xb4, 0x2d, 0xe9, 0x88, 0x24, 0x12, 0x8f,
	0xc6, 0x59, 0x56, 0xf7, 0xc3, 0xab, 0x84, 0xa9, 0xc0, 0xe3, 0x31, 0x11, 0x89, 0x8e, 0xdf, 0xab,
	0x32, 0x26, 0x6c, 0x32, 0xca, 0x09, 0x95, 0x95, 0x05, 0x9c, 0x49, 0x1c, 0x48, 0x4d, 0xe9, 0x55,
	0x51, 0x62, 0x1e, 0x60, 0x49, 0x39, 0xcb, 0x38, 0xbd, 0xb7, 0x06, 0xac, 0xfd, 0x90, 0x8d, 0xe7,
	0x15, 0x4d, 0xe8, 0x09, 0x8d, 0xa9, 0xbc, 0x40, 0x5b, 0x60, 0x47, 0x5c, 0x4c, 0xb1, 0x08, 0x89,
	0xf0, 0x19, 0x91, 0x3e, 0x0d, 0x1d, 0xd8, 0xac, 0x6d, 0xb5, 0xbc, 0x76, 0x81, 0x3f, 0x27, 0xf2,
	0x30, 0x44, 0x2e, 0x74, 0x4a, 0xa6, 0x24, 0x0c, 0x33, 0x45, 0xb6, 0x36, 0x6b, 0x5b, 0x4d, 0x6f,
	0xad, 0x08, 0x1d, 0xab, 0xc8, 0x61, 0x88, 0xbe, 0x80, 0xce, 0x29, 0x1f, 0x91, 0x54, 0x74, 0xca,
	0xc5, 0x59, 0x2e, 0xbe, 0xa2, 0xc4, 0xed, 0x34, 0xf4, 0x3c, 0x8b, 0x64, 0xf2, 0x3b, 0xb0, 0x71,
	0x89, 0x5e, 0x3a, 0xb4, 0x94, 0x43, 0x67, 0x26, 0xa3, 0xf0, 0x78, 0x0c, 0x30, 0x19, 0x87, 0x58,
	0x92, 0xd0, 0xc7, 0xd2, 0x69, 0x6f, 0xd6, 0xb6, 0xac, 0x7e, 0xd7, 0xcd, 0x26, 0xdc, 0xcd, 0x27,
	0xdc, 0x3d, 0xce, 0x57, 0xc4, 0x6b, 0x6a, 0xf6, 0x13, 0x89, 0xba, 0x60, 0xe6, 0x13, 0xe4, 0xd4,
	0x36, 0x6b, 0x5b, 0xa6, 0x57, 0xfc, 0xa3, 0xcf, 0x61, 0x0d, 0x33, 0x49, 0x18, 0xc3, 0xfe, 0x38,
	0xc6, 0x01, 0x19, 0x11, 0x26, 0x9d, 0x45, 0x45, 0xb2, 0x75, 0xe0, 0x45, 0x8e, 0xa3, 0x8f, 0xa0,
	0x95, 0x93, 0x03, 0x3e, 0x61, 0xd2, 0x59, 0x52, 0xc4, 0x15, 0x0d, 0xee, 0xa5, 0x18, 0x7a, 0x00,
	0xab, 0x11, 0x65, 0xc4, 0x2f, 0x36, 0x47, 0xe2, 0x18, 0x8a, 0xd6, 0x4e, 0xe1, 0xa2, 0xc0, 0x04,
	0xdd, 0x87, 0x15, 0xbd, 0xb4, 0x3e, 0x65, 0x11, 0x77, 0xea, 0x8a, 0x65, 0x69, 0xec, 0x90, 0x45,
	0x1c, 0x6d, 0x40, 0x23, 0x91, 0x58, 0x4e, 0x12, 0xa7, 0xa1, 0x82, 0xfa, 0x0f, 0x7d, 0x02, 0xed,
	0x48, 0x90, 0xd7, 0x13, 0xc2, 0x82, 0x8b, 0xb4, 0x6e, 0xe6, 0x98, 0x2a, 0xde, 0x2a, 0xd0, 0x17,
	0x31, 0x66, 0xa9, 0x43, 0xb6, 0x5b, 0x7c, 0x81, 0x25, 0x49, 0x9c, 0x66, 0xe6, 0x90, 0x61, 0x5e,
	0x0a, 0x0d, 0x0c, 0x73, 0xd9, 0x36, 0x7b, 0xff, 0x18, 0xb0, 0xae, 0x37, 0xcc, 0xfe, 0x25, 0x85,
	0x1d, 0x68, 0x08, 0x32, 0xcc, 0x27, 0xae, 0xdd, 0xbf, 0xed, 0x56, 0x1c, 0x0c, 0xd7, 0x53, 0x14,
	0x4f, 0x53, 0xd1, 0x6b, 0xd8, 0x88, 0xb9, 0xc0, 0xfe, 0x68, 0x12, 0x4b, 0xea, 0x27, 0x91, 0x1f,
	0x9c, 0x62, 0xc6, 0x48, 0x9c, 0x38, 0x4b, 0x9b, 0x4b, 0x5b, 0x56, 0xff, 0x9b, 0x4a, 0x91, 0x2a,
	0x7f, 0xf7, 0x29, 0xf7, 0xf0, 0xb3, 0x54, 0xe6, 0x68, 0x7f, 0x2f, 0x13, 0xf1, 0x3a, 0xa9, 0x76,
	0x86, 0x45, 0x1a, 0x4b, 0x90, 0x84, 0x0f, 0x94, 0x65, 0x42, 0xd9, 0x30, 0x26, 0x97, 0x3c, 0x0d,
	0xe5, 0xf9, 0xed, 0xfb, 0x79, 0x1e, 0x29, 0x9d, 0xd2, 0x74, 0x3d, 0x55, 0xd7, 0x60, 0xe9, 0xfa,
	0x12, 0xac, 0x28, 0x39, 0xcb, 0xad, 0xd4, 0x02, 0x5a, 0xfd, 0x47, 0xf3, 0x3b, 0xed, 0x1f, 0xfd,
	0x98, 0x1b, 0x40, 0x94, 0x9c, 0xe9, 0xef, 0x6e, 0x1f, 0xd0, 0xf5, 0x71, 0xa3, 0x3b, 0xd0, 0x2c,
	0x56, 0x57, 0xad, 0x86, 0xe1, 0x95, 0x40, 0xf7, 0x0f, 0xe8, 0x54, 0xd4, 0x7d, 0x73, 0x12, 0x7a,
	0x08, 0x76, 0x32, 0x16, 0x04, 0x87, 0x94, 0x0d, 0xfd, 0x08, 0x07, 0x92, 0x0b, 0xb5, 0xf7, 0x5b,
	0xde, 0x6a, 0x81, 0xef, 0x2b, 0x38, 0x15, 0x3a, 0xc1, 0x2c, 0x9c, 0xd2, 0x50, 0x9e, 0xaa, 0x6d,
	0xdf, 0xf2, 0x4a, 0xa0, 0xfb, 0x19, 0x40, 0x39, 0x96, 0x9b, 0x4d, 0x07, 0x86, 0xb9, 0x68, 0x2f,
	0xf5, 0xa6, 0x45, 0x87, 0x3a, 0x0c, 0x09, 0x93, 0x34, 0xa2, 0x44, 0x20, 0x17, 0x96, 0xc8, 0x84,
	0xaa, 0x14, 0xab, 0x7f, 0xe7, 0xda, 0xe1, 0x7e, 0x79, 0xc8, 0xe4, 0x57, 0x8f, 0x5e, 0xe1, 0x78,
	0x42, 0xbc, 0x94, 0x88, 0x36, 0xa0, 0x3e, 0x8e, 0x31, 0x65, 0xaa, 0xe8, 0xe6, 0xc1, 0x82, 0x97,
	0xfd, 0xa2, 0x75, 0x30, 0x4e, 0x71, 0x92, 0xd5, 0xb9, 0x72, 0xb0, 0xe0, 0xa9, 0xbf, 0x5d, 0x03,
	0x16, 0x69, 0xd8, 0xfb, 0xb3, 0x0e, 0xab, 0xda, 0xf9, 0x69, 0xde, 0x04, 0x8e, 0xc1, 0x92, 0x44,
	0x08, 0x92, 0x48, 0x41, 0x71, 0xac, 0xfd, 0xbf, 0xbc, 0x69, 0x1d, 0xf3, 0x54, 0xf7, 0xb8, 0xcc,
	0x3b, 0x58, 0xf0, 0x66, 0x65, 0xd0, 0x73, 0x68, 0x26, 0x58, 0x92, 0x38, 0xa6, 0x92, 0xa8, 0x0a,
	0xad, 0xbe, 0x3b, 0x97, 0xe6, 0x51, 0x9e, 0x75, 0xb0, 0xe0, 0x95, 0x12, 0xdd, 0xbf, 0x16, 0xc1,
	0x9a, 0xb1, 0x43, 0x8f, 0xaf, 0xb4, 0x35, 0xab, 0x7f, 0xb7, 0x52, 0x3e, 0xd7, 0x9d, 0xe9, 0x7a,
	0xbf, 0xbd, 0xab, 0xeb, 0xb5, 0xdf, 0x31, 0xec, 0x19, 0xdf, 0x27, 0x57, 0xba, 0x62, 0x45, 0x9f,
	0x7c, 0x52, 0xd5, 0x27, 0xdf, 0xb5, 0xa2, 0x3b, 0xfd, 0x6c, 0x45, 0x2f, 0x77, 0xd1, 0xbd, 0xea,
	0x2e, 0x5a, 0xd5, 0xf3, 0x77, 0x39, 0x8f, 0x33, 0x89, 0x2b, 0x1d, 0xb6, 0xfb, 0xb6, 0x06, 0xcd,
	0x62, 0x32, 0xff, 0xcf, 0x7c, 0xf5, 0xa0, 0x15, 0x51, 0x12, 0x87, 0x3e, 0x8f, 0xfc, 0x73, 0x4a,
	0xa6, 0x6a, 0xae, 0x6a, 0x9e, 0xa5, 0xc0, 0x9f, 0xa2, 0x57, 0x94, 0x4c, 0xd1, 0x2e, 0x80, 0x14,
	0xf8, 0x77, 0x92, 0x1e, 0x97, 0x0b, 0xdd, 0xe9, 0x7a, 0xd5, 0x93, 0x49, 0x47, 0x24, 0x2c, 0x5c,
	0x66, 0xb2, 0x76, 0x97, 0xa1, 0x7e, 0x9e, 0x8e, 0xa4, 0xf7, 0x4b, 0xd1, 0x8f, 0x73, 0x9e, 0x1a,
	0x21, 0xfa, 0xee, 0xda, 0x18, 0x3e, 0x9e, 0x67, 0x4b, 0x95, 0x43, 0xe9, 0xfd, 0x5d, 0x87, 0x65,
	0x1d, 0x45, 0xb7, 0xa0, 0xa1, 0xaf, 0xea, 0x9a, 0x3a, 0xd1, 0x75, 0xa6, 0xee, 0xe7, 0xdb, 0xd0,
	0x2c, 0xaf, 0xe4, 0xb6, 0xba, 0x92, 0x4d, 0x99, 0xdf, 0xc3, 0x77, 0x01, 0x82, 0x78, 0x92, 0x48,
	0x22, 0xd2, 0xa8, 0x3a, 0x78, 0x5e, 0x53, 0x23, 0x87, 0x21, 0xfa, 0x1e, 0x40, 0xbf, 0xa4, 0xd2,
	0xf0, 0xaa, 0x2a, 0xf1, 0xd3, 0x9b, 0x4a, 0x2c, 0x8f, 0xbf, 0xd7, 0x1c, 0xe6, 0xd0, 0x95, 0xdb,
	0xde, 0x78, 0x9f, 0xdb, 0x7e, 0x76, 0x8a, 0xd6, 0xff, 0xcb, 0x14, 0xa1, 0x9f, 0x61, 0x03, 0x87,
	0x23, 0xca, 0x68, 0x22, 0x05, 0x96, 0xf4, 0x9c, 0xf8, 0xfa, 0x4e, 0x76, 0x6c, 0xa5, 0xb7, 0x59,
	0xa9, 0xb7, 0x57, 0xde, 0xdb, 0xde, 0xad, 0xcb, 0xf9, 0x3a, 0x84, 0x9e, 0xc1, 0x9a, 0x24, 0xc1,
	0x29, 0xa3, 0x01, 0x8e, 0x0b, 0xcd, 0xb5, 0x39, 0x35, 0xed, 0x22, 0x35, 0x97, 0xeb, 0x43, 0x83,
	0xb3, 0x98, 0x32, 0xa2, 0x9e, 0x71, 0x37, 0x1f, 0x0d, 0xcd, 0x44, 0x2f, 0xae, 0xbd, 0x1c, 0x5a,
	0x2a, 0xf7, 0xe1, 0xdc, 0xb7, 0xd6, 0xd5, 0x47, 0xc6, 0x23, 0x58, 0x16, 0x6f, 0xd4, 0x03, 0xc3,
	0x41, 0x4a, 0xea, 0xf6, 0xb5, 0x32, 0xf6, 0x63, 0x8e, 0xa5, 0xae, 0x43, 0xbc, 0x49, 0x1f, 0x1e,
	0x69, 0x96, 0xd4, 0x59, 0x9d, 0x39, 0xb2, 0xa4, 0xca, 0x1a, 0x18, 0xe6, 0x92, 0x6d, 0x0c, 0x0c,
	0xb3, 0x6e, 0x37, 0x06, 0x86, 0xd9, 0xb0, 0x97, 0xb3, 0xf7, 0xcb, 0xc0, 0x30, 0x4d, 0xbb, 0x39,
	0x30, 0xcc, 0xa6, 0x0d, 0x03, 0xc3, 0xb4, 0xec, 0x95, 0x81, 0x61, 0xae, 0xd8, 0xad, 0xdd, 0x87,
	0xbf, 0x3e, 0x18, 0xf2, 0xcb, 0xa3, 0xe2, 0x62, 0xa8, 0x5f, 0xcd, 0x5f, 0xcf, 0xe2, 0x27, 0x0d,
	0xe5, 0xbb, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x38, 0xe5, 0x7e, 0x1d, 0x0c, 0x00,
	0x00,
}
