// Copyright © 2020 The Things Industries B.V.

package iampb

import (
	"errors"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *CreateNetworkRequest) Validate() error {
	if r.GetNetwork() == nil {
		return errors.New("network is required")
	}
	if r.Network.Authority != "" {
		return errors.New("custom authority is not allowed")
	}
	if err := packetbroker.NetID(r.Network.NetId).Validate(); err != nil {
		return err
	}
	for _, b := range r.Network.DevAddrBlocks {
		if err := b.Validate(); err != nil {
			return err
		}
	}
	for clusterID := range r.Network.NsIds {
		if !packetbroker.ClusterIDRegex.MatchString(clusterID) {
			return errors.New("invalid Cluster ID format")
		}
	}
	if r.Network.DelegatedNetId != nil {
		if err := packetbroker.NetID(r.Network.DelegatedNetId.Value).Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *NetworkRequest) Validate() error {
	return packetbroker.NetID(r.NetId).Validate()
}

// Validate returns whether the request is valid.
func (r *UpdateNetworkRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	for _, b := range r.GetDevAddrBlocks().GetValue() {
		if err := b.Validate(); err != nil {
			return err
		}
	}
	if nsIDs := r.NsIds; r != nil {
		for clusterID := range nsIDs.Value {
			if !packetbroker.ClusterIDRegex.MatchString(clusterID) {
				return errors.New("invalid Cluster ID format")
			}
		}
	}
	if delegatedNetID := r.DelegatedNetId.GetValue(); delegatedNetID != nil {
		if err := packetbroker.NetID(delegatedNetID.Value).Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateNetworkListedRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListTenantsRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateTenantRequest) Validate() error {
	if r.GetTenant() == nil {
		return errors.New("tenant is required")
	}
	if err := packetbroker.RequestTenantID(r.Tenant).Validate(); err != nil {
		return err
	}
	if r.Tenant.Authority != "" {
		return errors.New("custom authority is not allowed")
	}
	for _, b := range r.Tenant.DevAddrBlocks {
		if err := b.Validate(); err != nil {
			return err
		}
	}
	for clusterID := range r.Tenant.NsIds {
		if !packetbroker.ClusterIDRegex.MatchString(clusterID) {
			return errors.New("invalid Cluster ID format")
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *TenantRequest) Validate() error {
	return packetbroker.RequestTenantID(r).Validate()
}

// Validate returns whether the request is valid.
func (r *UpdateTenantRequest) Validate() error {
	if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
		return err
	}
	for _, b := range r.GetDevAddrBlocks().GetValue() {
		if err := b.Validate(); err != nil {
			return err
		}
	}
	if nsIDs := r.NsIds; r != nil {
		for clusterID := range nsIDs.Value {
			if !packetbroker.ClusterIDRegex.MatchString(clusterID) {
				return errors.New("invalid Cluster ID format")
			}
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateTenantListedRequest) Validate() error {
	if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
		return err
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListAPIKeysRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetClusterId()) {
		return errors.New("invalid Cluster ID format")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateAPIKeyRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetClusterId()) {
		return errors.New("invalid Cluster ID format")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *APIKeyRequest) Validate() error {
	if !packetbroker.APIKeyIDRegex.MatchString(r.GetKeyId()) {
		return errors.New("invalid API Key ID format")
	}
	return nil
}
