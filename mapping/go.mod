module go.packetbroker.org/api/mapping

go 1.15

replace go.packetbroker.org/api/v3 => ../v3

require (
	github.com/golang/protobuf v1.4.3
	go.packetbroker.org/api/v3 v3.1.1
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)
