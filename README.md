# go-unifi-network-server
A client for interacting with the Unifi Network Server (Controller) APIs.

> [!IMPORTANT]
> This is a work in progress and is subject to breaking changes without warning

The client currently supports the following endpoints:

- Client Devices

## Development

To generate an updated client run the following

```shell
go run internal/client-gen/main.go
```

Running integrations tests, which required Docker installed, can be done with the following:

```shell
go test ./...
```

## References And Inspiration

Below are links to places that have inspired or help construct the architecture of the client. Without the work of the
people on these projects this would not have been possible.

- https://github.com/paultyng/go-unifi
- https://github.com/unpoller/unifi
- https://github.com/google/go-github