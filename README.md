# zkproof-go

## Makefile targets

- Re-generate proto Go bindings: `make generate-proto`
- Build server: `make server`
- Build client: `make client`
- Clear binaries: `make clean`
- `make all` will clean, re-generate proto, test, build server and build client.
- `make test` will test ZK Go code.
- `make e2e` builds the server and client container images and run two test: happy and sad scenarios.
- `make server-img` builds the server img container.
- `make client-img` builds the client img container.

## zkproof-server

By default, the server is executed listening on `localhost:50051`.

```text
Usage of ./bin/zkproof-server:
  -host string
        connect to hostname (default "localhost")
  -port string
        TCP port (default "50051")
```

## zkproof-client

By default, the server is connects to `localhost:50051` as user `testUser` with password `1`.

```text
Usage of ./bin/zkproof-client:
  -host string
        connect to hostname (default "localhost")
  -password int
        Password (default 1)
  -port string
        TCP port (default "50051")
  -user string
        Username (default "testUser")
```

### Manual test

- `make all`
- Run `./bin/zkproof-server &`
- Run `./bin/zkproof-client`
- Run `killall zkproof-server` to stop the server.
