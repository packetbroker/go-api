// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var tenantIDRegexp = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// TenantID identifies a tenant of a Member.
type TenantID struct {
	NetID
	ID string
}

func (t TenantID) String() string {
	id := t.ID
	if id == "" {
		id = "default"
	}
	return fmt.Sprintf("%s/%s", t.NetID, id)
}

// IsEmpty indicates whether the TenantID is empty.
func (t TenantID) IsEmpty() bool {
	return t.NetID == 0 && t.ID == ""
}

// Validate returns an error if the tenant ID is invalid.
func (t TenantID) Validate() error {
	if !tenantIDRegexp.MatchString(t.ID) {
		return fmt.Errorf("invalid tenant ID format")
	}
	return nil
}

// ParseTenantID parses the TenantID formatted with String().
func ParseTenantID(s string) (TenantID, error) {
	p := strings.SplitN(s, "/", 2)
	if len(p) != 2 {
		return TenantID{}, errors.New("invalid tenant ID format")
	}
	res := TenantID{ID: p[1]}
	if err := res.NetID.UnmarshalText([]byte(p[0])); err != nil {
		return TenantID{}, err
	}
	return res, res.Validate()
}

// TenantRequest is a request with a Tenant ID.
type TenantRequest interface {
	GetNetId() uint32
	GetTenantId() string
}

// RequestTenantID returns the TenantID of the given request.
func RequestTenantID(req TenantRequest) TenantID {
	return TenantID{
		NetID: NetID(req.GetNetId()),
		ID:    req.GetTenantId(),
	}
}

// ForwarderTenantRequest is a request with a Forwarder Tenant ID.
type ForwarderTenantRequest interface {
	GetForwarderNetId() uint32
	GetForwarderTenantId() string
}

// ForwarderTenantID returns the TenantID of the given interface.
func ForwarderTenantID(req ForwarderTenantRequest) TenantID {
	return TenantID{
		NetID: NetID(req.GetForwarderNetId()),
		ID:    req.GetForwarderTenantId(),
	}
}

// HomeNetworkTenantRequest is a request with a Home Network Tenant ID.
type HomeNetworkTenantRequest interface {
	GetHomeNetworkNetId() uint32
	GetHomeNetworkTenantId() string
}

// HomeNetworkTenantID returns the TenantID of the given interface.
func HomeNetworkTenantID(req HomeNetworkTenantRequest) TenantID {
	return TenantID{
		NetID: NetID(req.GetHomeNetworkNetId()),
		ID:    req.GetHomeNetworkTenantId(),
	}
}
