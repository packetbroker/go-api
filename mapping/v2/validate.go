// Copyright © 2021 The Things Industries B.V.

package mappingpb

import (
	"errors"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *GetDefaultGatewayVisibilityRequest) Validate() error {
	if r.GetForwarderTenantId() != "" {
		return packetbroker.ForwarderTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkGatewayVisibilityRequest) Validate() error {
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if r.GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *SetGatewayVisibilityRequest) Validate() error {
	if r.GetVisibility().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetVisibility()).Validate(); err != nil {
			return err
		}
	}
	if r.GetVisibility().GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r.GetVisibility()).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateGatewayRequest) Validate() error {
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
