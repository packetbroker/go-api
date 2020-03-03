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

func netIDPtr(v NetID) *NetID { return &v }

func TestNetIDFromDevAddr(t *testing.T) {
	for _, tc := range []struct {
		DevAddr uint32
		NetID   *NetID
	}{
		{
			DevAddr: 0x2601abcd,
			NetID:   netIDPtr(0x000013),
		},
		{
			DevAddr: 0xe3fcabcd,
			NetID:   netIDPtr(0x6001fe),
		},
		{
			DevAddr: 0xfde00000,
			NetID:   netIDPtr(0xc07800),
		},
		{
			DevAddr: 0xffffabcd,
		},
	} {
		id, ok := NetIDFromDevAddr(tc.DevAddr)
		if tc.NetID != nil {
			if !ok || id != *tc.NetID {
				t.Fatalf("Expected DevAddr %08X NetID to equal %q (but it was %q)", tc.DevAddr, *tc.NetID, id)
			}
		} else if ok {
			t.Fatalf("Expected DevAddr %08X not to have a valid NetID", tc.DevAddr)
		}
	}
}
