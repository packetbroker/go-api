// SPDX-FileCopyrightText: Copyright 2020 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Package routingpb contains the Packet Broker Routing API v1 for Go.
package routingpb

import (
	"errors"
	"fmt"

	packetbroker "go.packetbroker.org/api/v3"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Validate returns whether the request is valid.
func (r *ListDefaultPoliciesRequest) Validate() error {
	if err := r.GetUpdatedSince().CheckValid(); err != nil && r.GetUpdatedSince() != nil {
		return status.Error(codes.InvalidArgument, "invalid updated_since")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetDefaultPolicyRequest) Validate() error {
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
func (r *ListHomeNetworkPoliciesRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if err := r.GetUpdatedSince().CheckValid(); err != nil && r.GetUpdatedSince() != nil {
		return status.Error(codes.InvalidArgument, "invalid updated_since")
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
func (r *GetHomeNetworkPolicyRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if err := packetbroker.NetID(r.GetHomeNetworkNetId()).Validate(); err != nil {
		return fmt.Errorf("home network NetID: %w", err)
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
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
func (r *SetPolicyRequest) Validate() error {
	if r.GetPolicy().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetPolicy()).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if r.GetPolicy().GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r.GetPolicy()).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
		return nil
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListEffectivePoliciesRequest) Validate() error {
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
func (r *ListNetworksWithPolicyRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return fmt.Errorf("tenant ID: %w", err)
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *PublishUplinkMessageRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if err := r.GetMessage().Validate(); err != nil {
		return fmt.Errorf("message: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *PublishDownlinkMessageRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if err := packetbroker.NetID(r.GetHomeNetworkNetId()).Validate(); err != nil {
		return fmt.Errorf("home network NetID: %w", err)
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
	}
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if err := r.GetMessage().Validate(); err != nil {
		return fmt.Errorf("message: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UplinkMessageDeliveryStateChangeRequest) Validate() error {
	if r.GetStateChange() == nil {
		return errors.New("state is required")
	}
	if err := r.GetStateChange().Validate(); err != nil {
		return fmt.Errorf("state change: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *DownlinkMessageDeliveryStateChangeRequest) Validate() error {
	if r.GetStateChange() == nil {
		return errors.New("state is required")
	}
	if err := r.GetStateChange().Validate(); err != nil {
		return fmt.Errorf("state change: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *SubscribeForwarderRequest) Validate() error {
	if err := packetbroker.NetID(r.GetForwarderNetId()).Validate(); err != nil {
		return fmt.Errorf("forwarder NetID: %w", err)
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
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
func (r *SubscribeHomeNetworkRequest) Validate() error {
	if err := packetbroker.NetID(r.GetHomeNetworkNetId()).Validate(); err != nil {
		return fmt.Errorf("home network NetID: %w", err)
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
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
func (r *RouteUplinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetMessage().GetForwarderClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetMessage()).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetMessage().GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r.GetMessage()).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
		return nil
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *RouteDownlinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetMessage().GetForwarderClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetMessage()).Validate(); err != nil {
			return fmt.Errorf("forwarder tenant ID: %w", err)
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetMessage().GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r.GetMessage()).Validate(); err != nil {
			return fmt.Errorf("home network tenant ID: %w", err)
		}
		return nil
	}
	return nil
}
