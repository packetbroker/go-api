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
	if !packetbroker.ClusterIDRegex.MatchString(r.GetClusterId()) {
		return errors.New("invalid Cluster ID format")
	}
	if r.GetTenantId() != "" {
		return packetbroker.RequestTenantID(r).Validate()
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
