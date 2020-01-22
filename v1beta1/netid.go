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
