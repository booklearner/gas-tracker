# gas-tracker

Tiny API to get (close to) real-time gas prices on the Ethereum network.

## Ethereum

- [Ethereum Gas Documentation](https://ethereum.org/en/developers/docs/gas/)
- [EIP-1559](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1559.md)
- [Ethereum API documentation](https://playground.open-rpc.org/?schemaUrl=https://raw.githubusercontent.com/ethereum/execution-apis/assembled-spec/openrpc.json&uiSchema%5BappBar%5D%5Bui:splitView%5D=false&uiSchema%5BappBar%5D%5Bui:input%5D=false&uiSchema%5BappBar%5D%5Bui:examplesDropdown%5D=false)
- [Cloudflare Ethereum Gateway](https://www.cloudflare.com/distributed-web-gateway/#ethereum-gateway)

## Usage

Use `make` to fetch dependencies and build the application.

```console
; make deps
; make build
; ./gas-tracker
```

Run the CLI (via `go run`):

```console
; make run
```

Run the server locally:

```console
; make server
```

Fetch dependencies and build the application without using `make`:

```console
; go mod download all  # fetch dependencies
; go build -o gas-tracker cmd/cli.go
; ./gas-tracker
```
