// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"regexp"
)

var (
	// ClusterIDRegex is the regular expression for validating cluster identifiers.
	ClusterIDRegex = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")
	// SubscriptionGroupRegexp is the regular expression for validating subscription groups.
	SubscriptionGroupRegexp = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")
	// APIKeyIDRegex is the regular expression for validating API key identifiers.
	APIKeyIDRegex = regexp.MustCompile("^[ABCDEFGHIJKLMNOPQRSTUVWXYZ234567]{16}$")
)

// Validate returns whether the request is valid.
func (r *GetTenantRequest) Validate() error {
	return RequestTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SetTenantRequest) Validate() error {
	if r.Tenant == nil {
		return errors.New("tenant is required")
	}
	return RequestTenantID(r.Tenant).Validate()
}

// Validate returns whether the request is valid.
func (r *DeleteTenantRequest) Validate() error {
	return RequestTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *GetDefaultRoutingPolicyRequest) Validate() error {
	return ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *ListHomeNetworkRoutingPoliciesRequest) Validate() error {
	return ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *GetHomeNetworkRoutingPolicyRequest) Validate() error {
	if err := ForwarderTenantID(r).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SetRoutingPolicyRequest) Validate() error {
	if r.Policy == nil {
		return errors.New("policy is required")
	}
	if err := ForwarderTenantID(r.Policy).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r.Policy).Validate()
}

// Validate returns whether the request is valid.
func (r *PublishUplinkMessageRequest) Validate() error {
	if !ClusterIDRegex.MatchString(r.ForwarderClusterId) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	return ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *PublishDownlinkMessageRequest) Validate() error {
	if !ClusterIDRegex.MatchString(r.ForwarderClusterId) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if err := ForwarderTenantID(r).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeForwarderRequest) Validate() error {
	if !ClusterIDRegex.MatchString(r.ForwarderClusterId) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if !SubscriptionGroupRegexp.MatchString(r.Group) {
		return errors.New("invalid subscription group format")
	}
	return ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeHomeNetworkRequest) Validate() error {
	if !SubscriptionGroupRegexp.MatchString(r.Group) {
		return errors.New("invalid subscription group format")
	}
	return HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *RouteUplinkMessageRequest) Validate() error {
	if r.Message == nil {
		return errors.New("message is required")
	}
	if err := ForwarderTenantID(r.Message).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r.Message).Validate()
}

// Validate returns whether the request is valid.
func (r *RouteDownlinkMessageRequest) Validate() error {
	if r.Message == nil {
		return errors.New("message is required")
	}
	if err := ForwarderTenantID(r.Message).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r.Message).Validate()
}

// Validate returns whether the DevAddrPrefix is valid.
func (pf *DevAddrPrefix) Validate() error {
	if pf.Length > 32 {
		return errors.New("length too long")
	}
	return nil
}

// Validate returns whether the DevAddrPrefix is valid.
func (b *DevAddrBlock) Validate() error {
	if !ClusterIDRegex.MatchString(b.HomeNetworkClusterId) {
		return errors.New("invalid cluster ID format")
	}
	return b.Prefix.Validate()
}
