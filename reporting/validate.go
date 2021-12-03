// Copyright © 2021 The Things Industries B.V.

package reportingpb

import (
	"errors"
	"fmt"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *GetRoutedMessagesRequest) Validate() error {
	if r.ForwarderNetId != nil && r.ForwarderTenantId != nil {
		id := packetbroker.TenantID{
			NetID: packetbroker.NetID(r.ForwarderNetId.Value),
			ID:    r.ForwarderTenantId.Value,
		}
		if err := id.Validate(); err != nil {
			return err
		}
	}
	if r.HomeNetworkNetId != nil && r.HomeNetworkTenantId != nil {
		id := packetbroker.TenantID{
			NetID: packetbroker.NetID(r.HomeNetworkNetId.Value),
			ID:    r.HomeNetworkTenantId.Value,
		}
		if err := id.Validate(); err != nil {
			return err
		}
	}
	if r.Time == nil {
		return errors.New("time is required")
	}
	if period := r.GetPeriod(); period != nil {
		if period.From == nil {
			return errors.New("from period is required")
		}
		if period.To == nil {
			return errors.New("to period is required")
		}
		for _, month := range []uint32{period.From.Month, period.To.Month} {
			if month < 1 || month > 12 {
				return fmt.Errorf("invalid month %d", month)
			}
		}
	}
	return nil
}
