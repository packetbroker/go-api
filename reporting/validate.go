// SPDX-FileCopyrightText: Copyright 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Package reportingpb contains the Packet Broker Reporting API v1 for Go.
package reportingpb

import (
	"errors"
	"fmt"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *GetRoutedMessagesRequest) Validate() error {
	if r.GetForwarderNetId() != nil {
		netID := packetbroker.NetID(r.GetForwarderNetId().GetValue())
		if err := netID.Validate(); err != nil {
			return fmt.Errorf("forwarder NetID: %w", err)
		}
		if r.GetForwarderTenantId() != nil {
			id := packetbroker.TenantID{
				NetID: netID,
				ID:    r.GetForwarderTenantId().GetValue(),
			}
			if err := id.Validate(); err != nil {
				return fmt.Errorf("forwarder tenant ID: %w", err)
			}
		}
	}
	if r.GetHomeNetworkNetId() != nil {
		netID := packetbroker.NetID(r.GetHomeNetworkNetId().GetValue())
		if err := netID.Validate(); err != nil {
			return fmt.Errorf("home network NetID: %w", err)
		}
		if r.GetHomeNetworkTenantId() != nil {
			id := packetbroker.TenantID{
				NetID: netID,
				ID:    r.GetHomeNetworkTenantId().GetValue(),
			}
			if err := id.Validate(); err != nil {
				return fmt.Errorf("home network tenant ID: %w", err)
			}
		}
	}
	if r.Time == nil {
		return errors.New("time is required")
	}
	if period := r.GetPeriod(); period != nil {
		if period.GetFrom() == nil {
			return errors.New("from period is required")
		}
		if period.GetTo() == nil {
			return errors.New("to period is required")
		}
		for _, month := range []uint32{period.GetFrom().GetMonth(), period.GetTo().GetMonth()} {
			if month < 1 || month > 12 {
				return fmt.Errorf("invalid month %d", month)
			}
		}
	}
	return nil
}
