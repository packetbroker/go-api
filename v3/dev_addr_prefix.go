// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MarshalText returns the DevAddrPrefix as formatted string.
func (m *DevAddrPrefix) MarshalText() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%08X/%d", m.Value, m.Length)), nil
}

// UnmarshalText parses the formatted string.
func (m *DevAddrPrefix) UnmarshalText(text []byte) error {
	parts := strings.SplitN(string(text), "/", 2)
	if len(parts) != 2 {
		return errors.New("packetbroker: require 2 parts in formatted DevAddrPrefix")
	}
	if len(parts[0]) != 8 {
		return errors.New("packetbroker: require hex value of length 8")
	}
	value, err := strconv.ParseUint(parts[0], 16, 32)
	if err != nil {
		return fmt.Errorf("packetbroker: invalid hex value: %w", err)
	}
	length, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return fmt.Errorf("packetbroker: invalid length value: %w", err)
	}
	*m = DevAddrPrefix{
		Value:  uint32(value),
		Length: uint32(length),
	}
	return nil
}

// Match returns true if the DevAddrPrefix matches the given DevAddr.
func (m *DevAddrPrefix) Match(devAddr uint32) bool {
	if m == nil {
		return false
	}
	shift := 32 - m.Length
	return m.Value>>shift == devAddr>>shift
}
