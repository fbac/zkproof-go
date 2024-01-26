# zkproof-go

## Makefile targets

- Re-generate proto Go bindings: `make generate-proto`
- Build server: `make server`
- Build client: `make client`
- Clear binaries: `make clean`
- `make all` will clean, re-generate proto, test, build server and build client.

## zkproof-server

By default, the server is executed listening on `localhost:50051`

```text
Usage of ./bin/zkproof-server:
  -host string
        connect to hostname (default "localhost")
  -port string
        TCP port (default "50051")
```

## zkproof-client

By default, the server is connects to `localhost:50051` as user `testUser` with password `1`

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
