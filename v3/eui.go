// Copyright © 2023 The Things Industries B.V.

package packetbroker

import "fmt"

// EUI represents a 64-bit Extended Unique Identifier.
type EUI uint64

// String implements Stringer.
func (n EUI) String() string {
	return fmt.Sprintf("%016X", uint64(n))
}

// MarshalText implements TextMarshaler.
func (n EUI) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

// UnmarshalText implements TextUnmarshaler.
func (n *EUI) UnmarshalText(buf []byte) error {
	_, err := fmt.Sscanf(string(buf), "%016X", n)
	return err
}
