// Copyright © 2021 The Things Industries B.V.

package mappingpb

import (
	"errors"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *GetDefaultGatewayVisibilityRequest) Validate() error {
	if err := packetbroker.NetID(r.ForwarderNetId).Validate(); err != nil {
		return err
	}
	if r.GetForwarderTenantId() != "" {
		return packetbroker.ForwarderTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkGatewayVisibilityRequest) Validate() error {
	if err := packetbroker.NetID(r.ForwarderNetId).Validate(); err != nil {
		return err
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if err := packetbroker.NetID(r.HomeNetworkNetId).Validate(); err != nil {
		return err
	}
	if r.GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *SetGatewayVisibilityRequest) Validate() error {
	if r.Visibility == nil {
		return errors.New("visibility is required")
	}
	if err := packetbroker.NetID(r.Visibility.ForwarderNetId).Validate(); err != nil {
		return err
	}
	if r.Visibility.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.Visibility).Validate(); err != nil {
			return err
		}
	}
	if err := packetbroker.NetID(r.Visibility.HomeNetworkNetId).Validate(); err != nil {
		return err
	}
	if r.Visibility.GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r.Visibility).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateGatewayRequest) Validate() error {
	if err := packetbroker.NetID(r.ForwarderNetId).Validate(); err != nil {
		return err
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderGatewayId() == nil {
		return errors.New("forwarder gateway ID is required")
	}
	if err := r.ForwarderGatewayId.Validate(); err != nil {
		return err
	}
	if r.Online.GetValue() && r.OnlineTtl == nil {
		return errors.New("online TTL is required when online is set and true")
	}
	return nil
}
