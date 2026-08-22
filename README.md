# NodeSentinel

Linux service monitoring laboratory written in Go.

It currently queries the `nginx` service through `systemctl` and displays its
load, active, and sub states.

## Requirements

- Go 1.27.0
- Linux with systemd
- `systemctl` available on `PATH`

## Usage

```sh
go run .
```

Example output:

```text
Service: nginx
Load State: loaded
Active State: active
Sub State: running
Exists: true
Running: true
```

The service name is currently hardcoded as `nginx` in `main.go`. If `systemctl`
fails, the application prints the error and exits with status 0.

## Verification

```sh
gofmt -w main.go
go vet ./...
go test ./...
go build ./...
```
