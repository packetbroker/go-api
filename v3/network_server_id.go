// Copyright © 2023 The Things Industries B.V.

package packetbroker

// NetworkServerID combines a LoRaWAN NetID and NSID to uniquely identify an instance of a Network Server.
type NetworkServerID struct {
	NetID NetID
	NSID  EUI
}
