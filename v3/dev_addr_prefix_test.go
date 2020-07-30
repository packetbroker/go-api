// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"bytes"
	"testing"
)

func TestDevAddrPrefix(t *testing.T) {
	p := &DevAddrPrefix{Value: 0x26011000, Length: 20}

	if v := p.String(); v != "26011000/20" {
		t.Fatalf("String() result %q does not equal %q", v, "26011000/20")
	}

	if !p.Match(0x26011abc) {
		t.Fatalf("%v should match %s", p, "26011ABC")
	}
	if p.Match(0x42424242) {
		t.Fatalf("%v should not match %s", p, "42424242")
	}

	buf, err := p.MarshalText()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(buf, []byte("26011000/20")) {
		t.Fatalf("MarshalText() result %v does not equal %v", buf, []byte("26011000/20"))
	}

	if err := p.UnmarshalText([]byte("27000000/8")); err != nil {
		t.Fatalf("UnmarshalText() %v failed: %v", []byte("27000000/8"), err)
	}
	expected := &DevAddrPrefix{Value: 0x27000000, Length: 8}
	if p.Value != expected.Value || p.Length != expected.Length {
		t.Fatalf("UnmarshalText() result %v does not equal %v", p, expected)
	}
}
