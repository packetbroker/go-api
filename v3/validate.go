// Copyright © 2020 The Things Industries B.V.

package packetbroker

import (
	"errors"
	"net/url"
	"regexp"
)

var (
	// ClusterIDRegex is the regular expression for validating cluster identifiers.
	ClusterIDRegex = regexp.MustCompile(`^(?:(?:[a-z0-9]|(?:[a-z0-9][a-z0-9-]?)*[a-z0-9])\.)*(?:[a-z0-9]|(?:[a-z0-9][a-z0-9-]?)*[a-z0-9])$|^$`)
	// SubscriptionGroupRegexp is the regular expression for validating subscription groups.
	SubscriptionGroupRegexp = regexp.MustCompile(`^(?:(?:[a-z0-9]|(?:[a-z0-9][a-z0-9-]?)*[a-z0-9])\.)*(?:[a-z0-9]|(?:[a-z0-9][a-z0-9-]?)*[a-z0-9])$|^$`)
	// APIKeyIDRegex is the regular expression for validating API key identifiers.
	APIKeyIDRegex = regexp.MustCompile("^[ABCDEFGHIJKLMNOPQRSTUVWXYZ234567]{16}$")
)

// Validate returns whether the DevAddrPrefix is valid.
func (pf *DevAddrPrefix) Validate() error {
	if pf.GetLength() > 32 {
		return errors.New("length too long")
	}
	return nil
}

// Validate returns whether the DevAddrPrefix is valid.
func (b *DevAddrBlock) Validate() error {
	if b.GetPrefix() == nil {
		return errors.New("prefix is required")
	}
	if !ClusterIDRegex.MatchString(b.GetHomeNetworkClusterId()) {
		return errors.New("invalid cluster ID format")
	}
	return b.Prefix.Validate()
}

// Validate returns whether the JoinEUIPrefix is valid.
func (pf *JoinEUIPrefix) Validate() error {
	if pf.GetLength() > 64 {
		return errors.New("length too long")
	}
	return nil
}

// Validate returns whether the Target is valid.
func (t *Target) Validate() error {
	switch t.Protocol {
	// LoRaWAN Backend Interfaces require a valid URL, or empty value for lookup.
	case Protocol_TS002_V1_0, Protocol_TS002_V1_1:
		if t.Address == "" {
			return nil
		}
		_, err := url.Parse(t.Address)
		return err
	default:
		return errors.New("invalid target protocol")
	}
}

func (e *JoinServerFixedEndpoint) Validate() error {
	if !ClusterIDRegex.MatchString(e.GetClusterId()) {
		return errors.New("invalid cluster ID format")
	}
	return nil
}

// Validate returns whether the GatewayIdentifier is valid.
func (i *GatewayIdentifier) Validate() error {
	if i.Eui == nil && len(i.GetPlain()) == 0 && len(i.GetHash()) == 0 {
		return errors.New("no gateway identifier specified")
	}
	if hash := i.GetHash(); hash != nil && len(hash) != 32 {
		return errors.New("invalid SHA-256 hash length")
	}
	return nil
}

func (m *UplinkMessage) Validate() error {
	if m.GetGatewayId() != nil {
		if err := m.GatewayId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (m *DownlinkMessage) Validate() error {
	return nil
}

// Validate returns whether the UplinkMessageDeliveryStateChange is valid.
func (c *UplinkMessageDeliveryStateChange) Validate() error {
	if !ClusterIDRegex.MatchString(c.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if c.GetForwarderTenantId() != "" {
		if err := ForwarderTenantID(c).Validate(); err != nil {
			return err
		}
	}
	if !ClusterIDRegex.MatchString(c.GetHomeNetworkClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if c.GetHomeNetworkTenantId() != "" {
		if err := HomeNetworkTenantID(c).Validate(); err != nil {
			return err
		}
	}
	if c.State != MessageDeliveryState_PROCESSED && c.Error != nil {
		return errors.New("error information set while delivery state is not PROCESSED")
	}
	return nil
}

// Validate returns whether the DownlinkMessageDeliveryStateChange is valid.
func (c *DownlinkMessageDeliveryStateChange) Validate() error {
	if !ClusterIDRegex.MatchString(c.GetHomeNetworkClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if c.GetHomeNetworkTenantId() != "" {
		if err := HomeNetworkTenantID(c).Validate(); err != nil {
			return err
		}
	}
	if !ClusterIDRegex.MatchString(c.GetForwarderClusterId()) {
		return errors.New("invalid Forwarder Cluster ID format")
	}
	if c.GetForwarderTenantId() != "" {
		if err := ForwarderTenantID(c).Validate(); err != nil {
			return err
		}
	}
	if id := c.GetForwarderGatewayId(); id != nil {
		if err := id.Validate(); err != nil {
			return err
		}
	}
	if c.State != MessageDeliveryState_PROCESSED {
		switch c.Result.(type) {
		case *DownlinkMessageDeliveryStateChange_Success:
			return errors.New("success informtion set while delivery state is not PROCESSED")
		case *DownlinkMessageDeliveryStateChange_Error:
			return errors.New("error information set while delivery state is not PROCESSED")
		}
	}
	return nil
}
