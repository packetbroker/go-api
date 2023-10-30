// Copyright © 2020 The Things Industries B.V.

package packetbroker

import "fmt"

// NetID represents a LoRa Alliance NetID.
type NetID uint32

// String implements Stringer.
func (n NetID) String() string {
	return fmt.Sprintf("%06X", uint32(n))
}

// MarshalText implements TextMarshaler.
func (n NetID) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

// UnmarshalText implements TextUnmarshaler.
func (n *NetID) UnmarshalText(buf []byte) error {
	_, err := fmt.Sscanf(string(buf), "%06X", n)
	return err
}

// Group returns the grouped NetID.
// For any Type 7 NetIDs, the first NetID in the group of 16 continuous Type 7 NetIDs is returned.
func (n NetID) Group() NetID {
	switch n >> 21 {
	case 0b111:
		return n & 0b111_11111111111111111_0000
	default:
		return n
	}
}

// DevAddrPrefix returns the DevAddr prefix.
func (n NetID) DevAddrPrefix(grouped bool) *DevAddrPrefix {
	switch n >> 21 {
	case 0b000: // Type 0
		return &DevAddrPrefix{Value: 0b0_000000_0000000000000000000000000 + uint32(n)&0b111111<<25, Length: 32 - 25}
	case 0b001: // Type 1
		return &DevAddrPrefix{Value: 0b10_000000_000000000000000000000000 + uint32(n)&0b111111<<24, Length: 32 - 24}
	case 0b010: // Type 2
		return &DevAddrPrefix{Value: 0b110_000000000_00000000000000000000 + uint32(n)&0b111111111<<20, Length: 32 - 20}
	case 0b011: // Type 3
		return &DevAddrPrefix{Value: 0b1110_00000000000_00000000000000000 + uint32(n)&0b11111111111<<17, Length: 32 - 17}
	case 0b100: // Type 4
		return &DevAddrPrefix{Value: 0b11110_000000000000_000000000000000 + uint32(n)&0b111111111111<<15, Length: 32 - 15}
	case 0b101: // Type 5
		return &DevAddrPrefix{Value: 0b111110_0000000000000_0000000000000 + uint32(n)&0b1111111111111<<13, Length: 32 - 13}
	case 0b110: // Type 6
		return &DevAddrPrefix{Value: 0b1111110_000000000000000_0000000000 + uint32(n)&0b111111111111111<<10, Length: 32 - 10}
	case 0b111: // Type 7
		if grouped {
			// LoRa Alliance issues groups of 16 continuous Type 7 NetIDs, aligned on 0xFFFFF0.
			return &DevAddrPrefix{Value: 0b11111110_00000000000000000_0000000 + uint32(n)&0b11111111111110000<<7, Length: 32 - 11}
		}
		return &DevAddrPrefix{Value: 0b11111110_00000000000000000_0000000 + uint32(n)&0b11111111111111111<<7, Length: 32 - 7}
	default:
		panic("unreachable")
	}
}

// MatchPrefix returns whether the given prefix falls within the NetID's DevAddr prefix.
// If grouped is true, the DevAddr prefix of the NetID group is considered. See NetID.Group().
func (n NetID) MatchPrefix(p *DevAddrPrefix, grouped bool) bool {
	netIDPrefix := n.DevAddrPrefix(grouped)
	if netIDPrefix.Length > p.Length {
		return false
	}
	outerL, outerH := netIDPrefix.Range()
	innerL, innerH := p.Range()
	return outerL <= innerL && outerH >= innerH
}

// NetIDFromDevAddr returns the sanitized NetID from the given DevAddr.
// See https://lora-alliance.org/sites/default/files/2019-08/tc22-00003.000.cr-be-type3-type4-netid-correction.pdf.
// If grouped is true, the first NetID of a group of NetIDs is returned. See NetID.Group().
func NetIDFromDevAddr(devAddr uint32, grouped bool) (NetID, bool) {
	switch {
	case devAddr&0x80000000 == 0x0: // NetID Type 0.
		return NetID(0x000000 + devAddr&0x7e000000>>25), true
	case devAddr&0xc0000000 == 0x80000000: // NetID Type 1.
		return NetID(0x200000 + devAddr&0x3f000000>>24), true
	case devAddr&0xe0000000 == 0xc0000000: // NetID Type 2.
		return NetID(0x400000 + devAddr&0x1ff00000>>20), true
	case devAddr&0xf0000000 == 0xe0000000: // NetID Type 3.
		return NetID(0x600000 + devAddr&0x0ffe0000>>17), true
	case devAddr&0xf8000000 == 0xf0000000: // NetID Type 4.
		return NetID(0x800000 + devAddr&0x07ff8000>>15), true
	case devAddr&0xfc000000 == 0xf8000000: // NetID Type 5.
		return NetID(0xa00000 + devAddr&0x03ffe000>>13), true
	case devAddr&0xfe000000 == 0xfc000000: // NetID Type 6.
		return NetID(0xc00000 + devAddr&0x01fffc00>>10), true
	case devAddr&0xff000000 == 0xfe000000: // NetID Type 7.
		if grouped {
			// NOTE: Return the first NetID in the group of 16 continuous Type 7 NetIDs.
			return NetID(0xe00000 + devAddr&0x00fff800>>7), true
		}
		return NetID(0xe00000 + devAddr&0x00ffff80>>7), true
	default:
		return 0, false
	}
}
