module go.packetbroker.org/api/mapping/v2

go 1.24.4

require (
	go.packetbroker.org/api/v3 v3.17.1
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.7
)

require (
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c // indirect
)

// Tagged releases are discontinued; depend on a master commit instead: go get go.packetbroker.org/api/mapping/v2@master
retract [v2.0.0, v2.3.3]
