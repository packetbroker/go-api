// SPDX-FileCopyrightText: Copyright 2020 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Package iampb contains the Packet Broker IAM API v1 for Go.
package iampb

import (
	"errors"
	"fmt"

	packetbroker "github.com/packetbroker/go-api/v3"
)

// Validate returns whether the request is valid.
func (r *CreateNetworkRequest) Validate() error {
	if r.GetNetwork() == nil {
		return errors.New("network is required")
	}
	if r.GetNetwork().GetAuthority() != "" {
		return errors.New("custom authority is not allowed")
	}
	if r.GetNetwork().GetNetId() == 0 {
		return errors.New("NetID is required")
	}
	if err := packetbroker.NetID(r.GetNetwork().GetNetId()).Validate(); err != nil {
		return fmt.Errorf("network NetID: %w", err)
	}
	if r.GetNetwork().GetDelegatedNetId() != nil {
		if err := packetbroker.NetID(r.GetNetwork().GetDelegatedNetId().GetValue()).Validate(); err != nil {
			return fmt.Errorf("delegated NetID: %w", err)
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *NetworkRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateNetworkRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	if delegatedNetID := r.GetDelegatedNetId().GetValue(); delegatedNetID != nil {
		if err := packetbroker.NetID(delegatedNetID.GetValue()).Validate(); err != nil {
			return fmt.Errorf("delegated NetID: %w", err)
		}
	}
	if target := r.GetTarget().GetValue(); target != nil {
		if err := target.Validate(); err != nil {
			return fmt.Errorf("target: %w", err)
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateNetworkListedRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListTenantsRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateTenantRequest) Validate() error {
	if r.GetTenant() == nil {
		return errors.New("tenant is required")
	}
	if err := packetbroker.RequestTenantID(r.GetTenant()).Validate(); err != nil {
		return fmt.Errorf("tenant ID: %w", err)
	}
	if r.GetTenant().GetAuthority() != "" {
		return errors.New("custom authority is not allowed")
	}
	for _, b := range r.GetTenant().GetDevAddrBlocks() {
		if err := b.Validate(); err != nil {
			return fmt.Errorf("DevAddr block: %w", err)
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *TenantRequest) Validate() error {
	if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
		return fmt.Errorf("tenant ID: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateTenantRequest) Validate() error {
	if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
		return fmt.Errorf("tenant ID: %w", err)
	}
	for _, b := range r.GetDevAddrBlocks().GetValue() {
		if err := b.Validate(); err != nil {
			return fmt.Errorf("DevAddr block: %w", err)
		}
	}
	if target := r.GetTarget().GetValue(); target != nil {
		if err := target.Validate(); err != nil {
			return fmt.Errorf("target: %w", err)
		}
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *UpdateTenantListedRequest) Validate() error {
	if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
		return fmt.Errorf("tenant ID: %w", err)
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *ListAPIKeysRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return fmt.Errorf("tenant ID: %w", err)
		}
	}
	if !packetbroker.ClusterIDRegex.MatchString(r.GetClusterId()) {
		return errors.New("invalid Cluster ID format")
	}
	return nil
}

// Validate returns whether the request is valid.
func (r *CreateAPIKeyRequest) Validate() error {
	if err := packetbroker.NetID(r.GetNetId()).Validate(); err != nil {
		return fmt.Errorf("NetID: %w", err)
	}
	if r.GetTenantId() != "" {
		if err := packetbroker.RequestTenantID(r).Validate(); err != nil {
			return fmt.Errorf("tenant ID: %w", err)
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
