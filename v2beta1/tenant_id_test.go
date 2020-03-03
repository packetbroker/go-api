// Copyright © 2020 The Things Industries B.V.

package packetbroker

import "testing"

func TestTenantID(t *testing.T) {
	for id, ok := range map[string]bool{
		"":             true,
		"valid":        true,
		"0invalid":     true,
		"12345":        true,
		"test-tenant":  true,
		"-test-tenant": false,
		"test-tenant-": false,
		"test--tenant": false,
		"INVALID":      false,
	} {
		t.Run(id, func(t *testing.T) {
			tenantID := TenantID{
				NetID: 0x0,
				ID:    id,
			}
			if (tenantID.Validate() == nil) != ok {
				t.Fatalf("Expected %q to be %s",
					id,
					map[bool]string{true: "valid", false: "invalid"}[ok],
				)
			}
		})
	}
}
