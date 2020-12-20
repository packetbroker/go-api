module go.packetbroker.org/api/mapping

go 1.15

replace go.packetbroker.org/api/v3 => ../v3

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/google/go-cmp v0.5.0 // indirect
	go.packetbroker.org/api/v3 v3.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.32.0
)
