// Copyright © 2020 The Things Industries B.V.

package routingpb

import (
	"errors"

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
	if r.GetForwarderTenantId() != "" {
		return packetbroker.ForwarderTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListHomeNetworkPoliciesRequest) Validate() error {
	if err := r.GetUpdatedSince().CheckValid(); err != nil && r.GetUpdatedSince() != nil {
		return status.Error(codes.InvalidArgument, "invalid updated_since")
	}
	if r.GetForwarderTenantId() != "" {
		return packetbroker.ForwarderTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkPolicyRequest) Validate() error {
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
func (r *SetPolicyRequest) Validate() error {
	if r.GetPolicy().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.GetPolicy()).Validate(); err != nil {
			return err
		}
	}
	if r.GetPolicy().GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r.GetPolicy()).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListEffectivePoliciesRequest) Validate() error {
	if r.GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListNetworksWithPolicyRequest) Validate() error {
	if r.GetTenantId() != "" {
		return packetbroker.RequestTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *PublishUplinkMessageRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if err := r.Message.Validate(); err != nil {
		return err
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *PublishDownlinkMessageRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if r.GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetHomeNetworkTenantId() != "" {
		if err := packetbroker.HomeNetworkTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if err := r.Message.Validate(); err != nil {
		return err
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UplinkMessageDeliveryStateChangeRequest) Validate() error {
	if r.GetStateChange() == nil {
		return errors.New("state is required")
	}
	return r.StateChange.Validate()
}

// Validate returns whether the request is valid.
func (r *DownlinkMessageDeliveryStateChangeRequest) Validate() error {
	if r.GetStateChange() == nil {
		return errors.New("state is required")
	}
	return r.StateChange.Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeForwarderRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
	}
	if r.GetForwarderTenantId() != "" {
		return packetbroker.ForwarderTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *SubscribeHomeNetworkRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
	}
	if r.GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *RouteUplinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.ForwarderClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.Message).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.HomeNetworkClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r.Message).Validate()
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *RouteDownlinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.ForwarderClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetForwarderTenantId() != "" {
		if err := packetbroker.ForwarderTenantID(r.Message).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.HomeNetworkClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if r.GetMessage().GetHomeNetworkTenantId() != "" {
		return packetbroker.HomeNetworkTenantID(r.Message).Validate()
	}
	return nil
}
