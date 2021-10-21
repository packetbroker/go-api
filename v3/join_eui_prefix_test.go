// Copyright © 2021 The Things Industries B.V.

package packetbroker

import (
	"bytes"
	"testing"
)

func TestJoinEUIPrefixes(t *testing.T) {
	p := &JoinEUIPrefix{Value: 0xec656e0000000000, Length: 24}

	if v, err := p.MarshalText(); err != nil {
		t.Fatalf("MarshalText() failed: %v", err)
	} else if string(v) != "EC656E0000000000/24" {
		t.Fatalf("MarshalText() result %q does not equal %q", v, "EC656E0000000000/24")
	}

	if !p.Match(0xec656e1000000000) {
		t.Fatalf("%v should match %s", p, "EC656E1000000000")
	}
	if p.Match(0x70b3d57ed0000001) {
		t.Fatalf("%v should not match %s", p, "70B3D57ED0000001")
	}

	buf, err := p.MarshalText()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(buf, []byte("EC656E0000000000/24")) {
		t.Fatalf("MarshalText() result %v does not equal %v", buf, []byte("EC656E0000000000/24"))
	}

	if err := p.UnmarshalText([]byte("70B3D57ED0000000/36")); err != nil {
		t.Fatalf("UnmarshalText() %v failed: %v", []byte("70B3D57ED0000000/36"), err)
	}
	expected := &JoinEUIPrefix{Value: 0x70b3d57ed0000000, Length: 36}
	if p.Value != expected.Value || p.Length != expected.Length {
		t.Fatalf("UnmarshalText() result %v does not equal %v", p, expected)
	}

	q := &JoinEUIPrefix{
		Value:  0xec656e0000000000,
		Length: 70,
	}
	if !q.Match(0xec656e0000000000) {
		t.Fatalf("%v should match %s even though the length is invalid", q, "EC656E0000000000")
	}
}
