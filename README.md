# Packet Broker API for Go

`go-api` is the Packet Broker API for Go.

## Installing

Use `go get` to retrieve the Go API.

```bash
$ go get go.packetbroker.org/api/v3
$ go get go.packetbroker.org/api/routing/v1
$ go get go.packetbroker.org/api/mapping/v2
$ go get go.packetbroker.org/api/iam
$ go get go.packetbroker.org/api/iam/v2
```

## Regenerating

The generated protos are checked in to this repository. You only need to regenerate the protos to incorporate changes coming from upstream.

To generate protos, [install the Protocol Compiler](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation).

Before generating protos, get the dependencies:

```bash
$ make deps
```

Clean existing generated protos:

```bash
$ make clean
```

To generate protos, clone the [Packet Broker API repository](https://github.com/packetbroker/api).

If you clone the API repository in `../../packetbroker/api` relative to the path of this repository:

```bash
$ make
```

If you clone the API repository somewhere else, pass the import path where the `packetbroker/api` folder resides like this:

```bash
$ PBAPI=<path> make
```

## License

The Go API is distributed under [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0). See `LICENSE` for more information.
