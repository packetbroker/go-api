// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"bytes"
	"testing"
)

func TestNetID(t *testing.T) {
	n := NetID(0x13)

	if v := n.String(); v != "000013" {
		t.Fatalf("String() result %q does not equal %q", v, "000013")
	}

	if p := (&DevAddrPrefix{Value: 0x26000000, Length: 7}); !n.MatchPrefix(p, false) {
		t.Fatalf("MatchPrefix() %v should match", p)
	}
	if p := (&DevAddrPrefix{Value: 0x26010000, Length: 16}); !n.MatchPrefix(p, false) {
		t.Fatalf("MatchPrefix() %v should match", p)
	}
	if p := (&DevAddrPrefix{Value: 0x26000000, Length: 6}); n.MatchPrefix(p, false) {
		t.Fatalf("MatchPrefix() %v should not match", p)
	}

	buf, err := n.MarshalText()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(buf, []byte("000013")) {
		t.Fatalf("MarshalText() result %v does not equal %v", buf, []byte("000013"))
	}

	if err := n.UnmarshalText([]byte("000042")); err != nil {
		t.Fatalf("UnmarshalText() %v failed: %v", []byte("000042"), err)
	}
	if n != 0x42 {
		t.Fatalf("UnmarshalText() result %v does not equal %v", uint32(n), 0x42)
	}
}

func TestDevAddrPrefixFromNetID(t *testing.T) {
	for _, tc := range []struct {
		NetID
		Prefix  *DevAddrPrefix
		Grouped bool
	}{
		{
			NetID:  0x000013,
			Prefix: &DevAddrPrefix{Value: 0x26000000, Length: 7},
		},
		{
			NetID:  0x600011,
			Prefix: &DevAddrPrefix{Value: 0xE0220000, Length: 15},
		},
		{
			NetID:  0xC0002F,
			Prefix: &DevAddrPrefix{Value: 0xFC00BC00, Length: 22},
		},
		{
			NetID:   0xE00060,
			Prefix:  &DevAddrPrefix{Value: 0xFE003000, Length: 21},
			Grouped: true,
		},
		{
			NetID:   0xE00061,
			Prefix:  &DevAddrPrefix{Value: 0xFE003000, Length: 21},
			Grouped: true,
		},
		{
			NetID:   0xE00060,
			Prefix:  &DevAddrPrefix{Value: 0xFE003000, Length: 25},
			Grouped: false,
		},
	} {
		if actual := tc.DevAddrPrefix(tc.Grouped); actual.Value != tc.Prefix.Value || actual.Length != tc.Prefix.Length {
			t.Fatalf("Expected NetID's %q DevAddr prefix to be %v (but it was %v)", tc.NetID, tc.Prefix, &actual)
		}
	}
}

func netIDPtr(v NetID) *NetID { return &v }

func TestNetIDFromDevAddr(t *testing.T) {
	for _, tc := range []struct {
		DevAddr uint32
		NetID   *NetID
		Grouped bool
	}{
		{
			DevAddr: 0x2601ABCD,
			NetID:   netIDPtr(0x000013),
		},
		{
			DevAddr: 0xE3FCABCD,
			NetID:   netIDPtr(0x6001FE),
		},
		{
			DevAddr: 0xFDE00000,
			NetID:   netIDPtr(0xC07800),
		},
		{
			DevAddr: 0xFE003000,
			NetID:   netIDPtr(0xE00060),
		},
		{
			DevAddr: 0xFE00378F, // Highest value in the highest block in the 16 Type 7 NetID blocks.
			NetID:   netIDPtr(0xE00060),
			Grouped: true,
		},
		{
			DevAddr: 0xFE00378F,
			NetID:   netIDPtr(0xE0006F),
			Grouped: false,
		},
		{
			DevAddr: 0xFFFFABCD,
		},
	} {
		id, ok := NetIDFromDevAddr(tc.DevAddr, tc.Grouped)
		if tc.NetID != nil {
			if !ok || id != *tc.NetID {
				t.Fatalf("Expected DevAddr %08X NetID to equal %q (but it was %q)", tc.DevAddr, *tc.NetID, id)
			}
		} else if ok {
			t.Fatalf("Expected DevAddr %08X not to have a valid NetID", tc.DevAddr)
		}
	}
}
