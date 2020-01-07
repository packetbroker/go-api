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
