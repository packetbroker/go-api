// SPDX-FileCopyrightText: Copyright 2020 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package packetbroker

import "testing"

func TestTenantID(t *testing.T) {
	t.Parallel()
	for input, ok := range map[string]bool{
		"valid":        true,
		"0invalid":     true,
		"12345":        true,
		"test-tenant":  true,
		"-test-tenant": false,
		"test-tenant-": false,
		"test--tenant": false,
		"INVALID":      false,
	} {
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			tenantID := TenantID{
				NetID: 0x13,
				ID:    input,
			}
			if (tenantID.Validate() == nil) != ok {
				t.Fatalf("Expected %q to be %s",
					input,
					map[bool]string{true: "valid", false: "invalid"}[ok],
				)
			}

			formatted := tenantID.String()
			parsed, err := ParseTenantID(formatted)
			if err != nil {
				if !ok {
					return
				}
				t.Fatalf("Failed to parse tenant ID: %s", err)
			} else if !ok {
				t.Fatalf("Expected parse invalid tenant ID to fail")
			}
			if tenantID.ID == "" {
				tenantID.ID = "default"
			}
			if tenantID != parsed {
				t.Fatal("Expected parsed tenant ID to equal")
			}
		})
	}
}
