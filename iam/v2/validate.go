// Copyright © 2021 The Things Industries B.V.

package iampb

import (
	"errors"

	packetbroker "go.packetbroker.org/api/v3"
)

// Validate returns whether the request is valid.
func (r *ListNetworkAPIKeysRequest) Validate() error {
	if id := r.GetClusterId(); id != nil {
		if !packetbroker.ClusterIDRegex.MatchString(id.GetValue()) {
			return errors.New("invalid Cluster ID format")
		}
	}
	if nid := r.GetNetId(); nid != nil {
		if err := packetbroker.NetID(nid.GetValue()).Validate(); err != nil {
			return err
		}
		if tid := r.GetTenantId(); tid != nil {
			tenantID := packetbroker.TenantID{
				NetID: packetbroker.NetID(nid.GetValue()),
				ID:    tid.GetValue(),
			}
			if err := tenantID.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateNetworkAPIKeyRequest) Validate() error {
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

// Validate returns whether the request is valid.
func (r *ListClusterAPIKeysRequest) Validate() error {
	if id := r.GetClusterId(); id != nil {
		if !packetbroker.ClusterIDRegex.MatchString(id.GetValue()) {
			return errors.New("invalid Cluster ID format")
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateClusterAPIKeyRequest) Validate() error {
	if r.GetClusterId() == "" {
		return errors.New("cluster ID is required")
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetClusterId()) {
		return errors.New("invalid Cluster ID format")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListNetworksRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return err
		}
	}
	if ref := r.GetPolicyReference(); ref != nil {
		if err := packetbroker.NetID(ref.NetId).Validate(); err != nil {
			return err
		}
		if ref.TenantId != "" {
			if err := packetbroker.RequestTenantID(r.PolicyReference).Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListJoinServersRequest) Validate() error {
	if err := packetbroker.NetID(r.NetId).Validate(); err != nil {
		return err
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateJoinServerRequest) Validate() error {
	if r.GetJoinServer() == nil {
		return errors.New("network is required")
	}
	if r.JoinServer.Id != 0 {
		return errors.New("ID cannot be specified")
	}
	for _, p := range r.JoinServer.JoinEuiPrefixes {
		if err := p.Validate(); err != nil {
			return err
		}
	}
	switch resolver := r.JoinServer.Resolver.(type) {
	case *packetbroker.JoinServer_Lookup:
		if err := resolver.Lookup.Validate(); err != nil {
			return err
		}
	case *packetbroker.JoinServer_Fixed:
		if err := resolver.Fixed.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateJoinServerRequest) Validate() error {
	for _, p := range r.GetJoinEuiPrefixes().GetValue() {
		if err := p.Validate(); err != nil {
			return err
		}
	}
	switch resolver := r.GetResolver().(type) {
	case *UpdateJoinServerRequest_Lookup:
		if err := resolver.Lookup.Validate(); err != nil {
			return err
		}
	case *UpdateJoinServerRequest_Fixed:
		if err := resolver.Fixed.Validate(); err != nil {
			return err
		}
	}
	return nil
}
