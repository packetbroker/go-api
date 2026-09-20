module go.packetbroker.org/api/iam

go 1.27.0

require (
	go.packetbroker.org/api/v3 v3.0.0-20260918125712-30f2e397fe07
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
)

require (
	golang.org/x/net v0.59.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260917231906-eeb232e0883d // indirect
)

// The go.packetbroker.org/api modules are frozen; development continues at github.com/packetbroker/go-api.
retract [v1.0.0, v1.8.3]
