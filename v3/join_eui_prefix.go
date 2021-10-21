// Copyright © 2021 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MarshalText returns the JoinEUIPrefix as formatted string.
func (m *JoinEUIPrefix) MarshalText() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%016X/%d", m.Value, m.Length)), nil
}

// UnmarshalText parses the formatted string.
func (m *JoinEUIPrefix) UnmarshalText(text []byte) error {
	parts := strings.SplitN(string(text), "/", 2)
	if len(parts) != 2 {
		return errors.New("packetbroker: require 2 parts in formatted JoinEUIPrefix")
	}
	if len(parts[0]) != 16 {
		return errors.New("packetbroker: require hex value of length 16")
	}
	value, err := strconv.ParseUint(parts[0], 16, 64)
	if err != nil {
		return fmt.Errorf("packetbroker: invalid hex value: %w", err)
	}
	length, err := strconv.ParseUint(parts[1], 10, 8)
	if err != nil {
		return fmt.Errorf("packetbroker: invalid length value: %w", err)
	}
	*m = JoinEUIPrefix{
		Value:  uint64(value),
		Length: uint32(length),
	}
	return nil
}

// Match returns true if the JoinEUIPrefix matches the given DevAddr.
func (m *JoinEUIPrefix) Match(joinEUI uint64) bool {
	if m == nil {
		return false
	}
	shift := 64 - m.Length
	return m.Value>>shift == joinEUI>>shift
}
