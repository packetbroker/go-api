// Copyright © 2021 The Things Industries B.V.

package packetbroker

// NewLoRaDataRate returns a LoRa data rate.
func NewLoRaDataRate(sf, bw uint32, codingRate string) *DataRate {
	return &DataRate{
		Modulation: &DataRate_Lora{
			Lora: &LoRaDataRate{
				SpreadingFactor: sf,
				Bandwidth:       bw,
				CodingRate:      codingRate,
			},
		},
	}
}

// NewFSKDataRate returns an FSK data rate.
func NewFSKDataRate(bps uint32) *DataRate {
	return &DataRate{
		Modulation: &DataRate_Fsk{
			Fsk: &FSKDataRate{
				BitsPerSecond: bps,
			},
		},
	}
}

// NewLRFHSSDataRate returns an LRFHSS data rate.
func NewLRFHSSDataRate(mt, ocw uint32, codr string) *DataRate {
	return &DataRate{
		Modulation: &DataRate_Lrfhss{
			Lrfhss: &LRFHSSDataRate{
				ModulationType:        mt,
				OperatingChannelWidth: ocw,
				CodingRate:            codr,
			},
		},
	}
}
