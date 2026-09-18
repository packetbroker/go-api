// SPDX-FileCopyrightText: Copyright 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package packetbroker

import (
	"bytes"
	"testing"
)

func TestJoinEUIPrefixes(t *testing.T) {
	t.Parallel()
	prefix := &JoinEUIPrefix{Value: 0xec656e0000000000, Length: 24}

	if v, err := prefix.MarshalText(); err != nil {
		t.Fatalf("MarshalText() failed: %v", err)
	} else if string(v) != "EC656E0000000000/24" {
		t.Fatalf("MarshalText() result %q does not equal %q", v, "EC656E0000000000/24")
	}

	if !prefix.Match(0xec656e1000000000) {
		t.Fatalf("%v should match %s", prefix, "EC656E1000000000")
	}
	if prefix.Match(0x70b3d57ed0000001) {
		t.Fatalf("%v should not match %s", prefix, "70B3D57ED0000001")
	}

	buf, err := prefix.MarshalText()
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(buf, []byte("EC656E0000000000/24")) {
		t.Fatalf("MarshalText() result %v does not equal %v", buf, []byte("EC656E0000000000/24"))
	}

	if err := prefix.UnmarshalText([]byte("70B3D57ED0000000/36")); err != nil {
		t.Fatalf("UnmarshalText() %v failed: %v", []byte("70B3D57ED0000000/36"), err)
	}
	expected := &JoinEUIPrefix{Value: 0x70b3d57ed0000000, Length: 36}
	if prefix.GetValue() != expected.GetValue() || prefix.GetLength() != expected.GetLength() {
		t.Fatalf("UnmarshalText() result %v does not equal %v", prefix, expected)
	}

	q := &JoinEUIPrefix{
		Value:  0xec656e0000000000,
		Length: 70,
	}
	if !q.Match(0xec656e0000000000) {
		t.Fatalf("%v should match %s even though the length is invalid", q, "EC656E0000000000")
	}
}
