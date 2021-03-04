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
