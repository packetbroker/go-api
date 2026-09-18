# Packet Broker API for Go

`go-api` is the Packet Broker API for Go.

## Installing

The Go API is split into Go modules, one per API. Tagged releases are discontinued: depend on a commit of the `master` branch instead (a pseudo-version). Go resolves the tip of `master` when it bypasses the module proxy for `go.packetbroker.org`:

```bash
$ export GOPRIVATE=go.packetbroker.org/*
$ go get go.packetbroker.org/api/v3@master
$ go get go.packetbroker.org/api/routing@master
$ go get go.packetbroker.org/api/routing/v2@master
$ go get go.packetbroker.org/api/mapping/v2@master
$ go get go.packetbroker.org/api/iam@master
$ go get go.packetbroker.org/api/iam/v2@master
$ go get go.packetbroker.org/api/reporting@master
```

Once a module is required by pseudo-version, `go get -u ./...` (with `GOPRIVATE` set as above) upgrades it to the tip of `master`. Through the public module proxy, `@latest` still resolves to the last (retracted) tag, so always use `@master` explicitly there.

## Regenerating

The generated code is checked in to this repository. You only need to regenerate the code to incorporate changes coming from the [Packet Broker API repository](https://github.com/packetbroker/api).

The protoc plugins are Go tool dependencies of the root module and need no installation. Install the [Protocol Compiler](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) at the version declared as `PROTOC_VERSION` in the `Makefile`; the generated code embeds the protoc version, so CI regenerates with exactly that version and fails on differences.

Clone the Packet Broker API repository such that its `packetbroker/api` folder is at `../../packetbroker/api` relative to this repository, check out the commit declared as `PBAPI_REF` in the `Makefile`, and run:

```bash
$ make clean all
```

If the API repository is cloned elsewhere, pass the directory that contains the `packetbroker/api` folder:

```bash
$ PBAPI=<path> make clean all
```

After regenerating, update `PBAPI_REF` in the `Makefile` to the API commit that the code was generated from.

## Development

The repository carries a `go.work` so that all modules build against the local `v3` module:

```bash
$ make build    # build every module standalone, against its published requirements
$ make test     # test every module standalone
$ go test ./... # test every module in the workspace, against the local v3 module
$ make quality  # lint all modules
$ make fmt      # format all modules
```

The service modules (`routing`, `iam`, `mapping/v2`, `reporting`, ...) require `go.packetbroker.org/api/v3` by pseudo-version. When a change touches `v3` and a service module at the same time, first commit and push the change, then update the service modules to that commit:

```bash
$ cd routing && GOPRIVATE=go.packetbroker.org/* go get go.packetbroker.org/api/v3@<commit> && go mod tidy
```

## License

The Go API is distributed under [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0). See `LICENSE` for more information.
