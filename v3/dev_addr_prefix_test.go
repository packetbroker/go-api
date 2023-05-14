// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"testing"
)

func TestDevAddrPrefix(t *testing.T) {
	p := &DevAddrPrefix{Value: 0x26011000, Length: 20}

	if !p.Match(0x26011abc) {
		t.Fatalf("%v should match %s", p, "26011ABC")
	}
	if p.Match(0x42424242) {
		t.Fatalf("%v should not match %s", p, "42424242")
	}

	low, high := p.Range()
	if low != 0x26011000 {
		t.Fatalf("%v should have lower bound %s", p, "26011000")
	}
	if high != 0x26011FFF {
		t.Fatalf("%v should have higher bound %s", p, "26011FFF")
	}

	if v, err := p.MarshalText(); err != nil {
		t.Fatalf("MarshalText() failed: %v", err)
	} else if string(v) != "26011000/20" {
		t.Fatalf("MarshalText() result %q does not equal %q", v, "26011000/20")
	}

	q := new(DevAddrPrefix)
	if err := q.UnmarshalText([]byte("27000000/8")); err != nil {
		t.Fatalf("UnmarshalText() %v failed: %v", []byte("27000000/8"), err)
	}
	expected := &DevAddrPrefix{Value: 0x27000000, Length: 8}
	if q.Value != expected.Value || q.Length != expected.Length {
		t.Fatalf("UnmarshalText() result %v does not equal %v", q, expected)
	}

	r := &DevAddrPrefix{
		Value:  0x27000000,
		Length: 34,
	}
	if !r.Match(0x27000000) {
		t.Fatalf("%v should match %s even though the length is invalid", r, "27000000")
	}
}
