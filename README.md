# Packet Broker API for Go

`go-api` is the Packet Broker API for Go, generated from the [Packet Broker API](https://github.com/packetbroker/api).

## Installing

The API is one Go module, `github.com/packetbroker/go-api`, with a package per API:

| Package | API |
| --- | --- |
| `github.com/packetbroker/go-api/v3` | Common types (`org.packetbroker.v3`) |
| `github.com/packetbroker/go-api/routing` | Routing v1 |
| `github.com/packetbroker/go-api/routing/v2` | Routing v2 |
| `github.com/packetbroker/go-api/iam` | IAM v1 |
| `github.com/packetbroker/go-api/iam/v2` | IAM v2 |
| `github.com/packetbroker/go-api/mapping/v2` | Mapping v2 |
| `github.com/packetbroker/go-api/reporting` | Reporting v1 |

There are no tagged releases: depend on a commit of the `master` branch (a pseudo-version), and `go get -u` follows `master` from there:

```bash
$ go get github.com/packetbroker/go-api@master
```

### Previous import path

The API was previously published as seven modules under `go.packetbroker.org/api/...`, one per package above. That import path is frozen: existing requirements keep resolving to the old commits and tags, but they receive no updates. To upgrade, replace the import path prefix `go.packetbroker.org/api/` with `github.com/packetbroker/go-api/`, require `github.com/packetbroker/go-api` and drop the seven old requirements.

## Regenerating

The generated code is checked in to this repository. You only need to regenerate the code to incorporate changes coming from the [Packet Broker API repository](https://github.com/packetbroker/api).

Code is generated with [buf](https://buf.build/). buf and the protoc plugins (`protoc-gen-go`, `protoc-gen-go-grpc`) are Go tool dependencies of the `tools` module, so that they do not become dependencies of the API itself, and need no installation. `buf.gen.yaml` maps each API to its Go package with buf's managed mode; the API definitions themselves carry no Go options.

By default, buf generates from the commit of the API repository declared as `PBAPI_REF` in the `Makefile`, which buf fetches itself:

```bash
$ make clean all
```

CI does the same and fails when the checked-in code differs, so the code and `PBAPI_REF` always match. To pick up API changes, set `PBAPI_REF` to the new API commit and regenerate.

To generate from a local checkout of the API repository instead, for instance to try unmerged API changes, point `PBAPI_INPUT` at it (the directory that contains `buf.yaml` and `packetbroker/api`):

```bash
$ PBAPI_INPUT=../api make clean all
```

Once the API changes are merged, set `PBAPI_REF` to the merge commit and run `make clean all` again before committing, so that the checked-in code is what CI regenerates.

## Development

```bash
$ make build      # build
$ make test       # test
$ make quality    # lint
$ make fmt        # format
$ make deps.tidy  # tidy the API and tools modules
```

## License

The Go API is distributed under [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0). See `LICENSE` for more information.
