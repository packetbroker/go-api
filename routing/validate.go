// Copyright © 2020 The Things Industries B.V.

package routingpb

import (
	"errors"

	pbtypes "github.com/gogo/protobuf/types"
	packetbroker "go.packetbroker.org/api/v3"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Validate returns whether the request is valid.
func (r *ListDefaultPoliciesRequest) Validate() error {
	if _, err := pbtypes.TimestampFromProto(r.GetUpdatedSince()); err != nil && r.GetUpdatedSince() != nil {
		return status.Error(codes.InvalidArgument, "invalid updated_since")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *GetDefaultPolicyRequest) Validate() error {
	return packetbroker.ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *ListHomeNetworkPoliciesRequest) Validate() error {
	if _, err := pbtypes.TimestampFromProto(r.GetUpdatedSince()); err != nil && r.GetUpdatedSince() != nil {
		return status.Error(codes.InvalidArgument, "invalid updated_since")
	}
	return packetbroker.ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkPolicyRequest) Validate() error {
	if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
		return err
	}
	return packetbroker.HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SetPolicyRequest) Validate() error {
	if err := packetbroker.ForwarderTenantID(r.GetPolicy()).Validate(); err != nil {
		return err
	}
	return packetbroker.HomeNetworkTenantID(r.GetPolicy()).Validate()
}

// Validate returns whether the request is valid.
func (r *PublishUplinkMessageRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	return packetbroker.ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *PublishDownlinkMessageRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if err := packetbroker.ForwarderTenantID(r).Validate(); err != nil {
		return err
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	return packetbroker.HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeForwarderRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
	}
	return packetbroker.ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeHomeNetworkRequest) Validate() error {
	if !packetbroker.ClusterIDRegex.MatchString(r.GetHomeNetworkClusterId()) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if !packetbroker.SubscriptionGroupRegexp.MatchString(r.GetGroup()) {
		return errors.New("invalid subscription group format")
	}
	return packetbroker.HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *RouteUplinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.ForwarderClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if err := packetbroker.ForwarderTenantID(r.Message).Validate(); err != nil {
		return err
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.HomeNetworkClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	return packetbroker.HomeNetworkTenantID(r.Message).Validate()
}

// Validate returns whether the request is valid.
func (r *RouteDownlinkMessageRequest) Validate() error {
	if r.GetMessage() == nil {
		return errors.New("message is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.ForwarderClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	if err := packetbroker.ForwarderTenantID(r.Message).Validate(); err != nil {
		return err
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.Message.HomeNetworkClusterId) {
		return errors.New("invalid Home Network Cluster ID format")
	}
	return packetbroker.HomeNetworkTenantID(r.Message).Validate()
}
