// SPDX-FileCopyrightText: Copyright 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Package mappingpb contains the Packet Broker Mapping API v2 for Go.
package mappingpb

import (
	"errors"
	"fmt"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *GetDefaultGatewayVisibilityRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
		return nil
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkGatewayVisibilityRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if err := packetbroker.NetID(r.GetHomeNetworkNetId()).Validate(); err != nil {
		return fmt.Errorf("home network NetID: %w", err)
	}
	if r.GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
		return nil
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *SetGatewayVisibilityRequest) Validate() error {
	if r.GetVisibility() == nil {
		return errors.New("visibility is required")
	}
	if err := packetbroker.NetID(r.GetVisibility().GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if r.GetVisibility().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetVisibility()).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if err := packetbroker.NetID(r.GetVisibility().GetHomeNetworkNetId()).Validate(); err != nil {
		return fmt.Errorf("home network NetID: %w", err)
	}
	if r.GetVisibility().GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r.GetVisibility()).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
		return nil
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateGatewayRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderGatewayId() == nil {
		return errors.New("forwarder gateway ID is required")
	}
	if err := r.GetForwarderGatewayId().Validate(); err != nil {
		return fmt.Errorf("forwarder gateway ID: %w", err)
	}
	if r.GetOnline().GetValue() && r.GetOnlineTtl() == nil {
		return errors.New("online TTL is required when online is set and true")
	}
	return nil
}
