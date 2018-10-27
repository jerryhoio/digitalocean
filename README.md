Jerryho DigitalOcean service
============================

TODO: Description

# Usage

Start a service

```
MICRO_REGISTRY=mdns go run main.go
```

Connect to the service

```
MICRO_REGISTRY=mdns micro call greeter Greeter.Hello '{"name": "John"}'
```

# TODO

1. Decide if `proto` directory is ok or not?

    https://github.com/golang-standards/project-layout/ has no `proto` directory.
    I assume we could copy Helm's layout which uses `_proto` for `*.proto` files and `pkg/proto` for compiled client code.
    I'm open for discussion. For now, leaving `proto` and `pkg/proto` as it's easier to script it this way.

[![Go Report Card](https://goreportcard.com/badge/jerryhoio/digitalocean)](https://goreportcard.com/report/jerryhoio/digitalocean)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/jerryhoio/digitalocean)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/jerryhoio/digitalocean/releases/latest)
