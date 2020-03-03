// Copyright © 2020 The Things Industries B.V.

package packetbroker

import "errors"

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
	return ForwarderTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *PublishDownlinkMessageRequest) Validate() error {
	if err := ForwarderTenantID(r).Validate(); err != nil {
		return err
	}
	return HomeNetworkTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *SubscribeForwarderRequest) Validate() error { return nil }

// Validate returns whether the request is valid.
func (r *SubscribeHomeNetworkRequest) Validate() error { return nil }

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
