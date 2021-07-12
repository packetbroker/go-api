// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"fmt"
)

// Endpoint defines an endpoint.
type Endpoint struct {
	TenantID
	ClusterID string
}

func (e Endpoint) String() string {
	clusterID := e.ClusterID
	if clusterID == "" {
		clusterID = "default"
	}
	return fmt.Sprintf("%s/%s", e.TenantID, clusterID)
}

// IsEmpty indicates whether the Endpoint is empty.
func (e Endpoint) IsEmpty() bool {
	return e.TenantID.IsEmpty() && e.ClusterID == ""
}

// Validate returns an error if the Endpoint is invalid.
func (e Endpoint) Validate() error {
	if err := e.TenantID.Validate(); err != nil {
		return err
	}
	if !ClusterIDRegex.MatchString(e.ClusterID) {
		return errors.New("invalid cluster ID format")
	}
	return nil
}

// ForwarderEndpointRequest is a request with a Forwarder Endpoint.
type ForwarderEndpointRequest interface {
	GetForwarderNetId() uint32
	GetForwarderTenantId() string
	GetForwarderClusterId() string
}

// ForwarderEndpoint returns the Endpoint of the given interface.
func ForwarderEndpoint(req ForwarderEndpointRequest) Endpoint {
	return Endpoint{
		TenantID:  ForwarderTenantID(req),
		ClusterID: req.GetForwarderClusterId(),
	}
}

// HomeNetworkEndpointRequest is a request with a Home Network Endpoint.
type HomeNetworkEndpointRequest interface {
	GetHomeNetworkNetId() uint32
	GetHomeNetworkTenantId() string
	GetHomeNetworkClusterId() string
}

// HomeNetworkEndpoint returns the Endpoint of the given interface.
func HomeNetworkEndpoint(req HomeNetworkEndpointRequest) Endpoint {
	return Endpoint{
		TenantID:  HomeNetworkTenantID(req),
		ClusterID: req.GetHomeNetworkClusterId(),
	}
}
