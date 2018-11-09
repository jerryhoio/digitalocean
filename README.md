Jerryho DigitalOcean service
============================

[![Go Report Card](https://goreportcard.com/badge/jerryhoio/digitalocean)](https://goreportcard.com/report/jerryhoio/digitalocean)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/jerryhoio/digitalocean)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/jerryhoio/digitalocean/releases/latest)

TODO: Description

# ROE

## PROTO naming

1. Put all `.proto` files in `proto/digitalocean`
2. Name `.proto` file after `digitalocean/godo` API structs, ex. `Droplets` => `droplets.proto`

# Usage

Start a service

```
MICRO_REGISTRY=mdns go run main.go
```

Connect to the service

```
MICRO_REGISTRY=mdns micro call greeter Greeter.Hello '{"name": "John"}'
```

#EXAMPLES:

1. List

```
MICRO_REGISTRY=mdns micro call digitalocean Droplets.List
```

2. Create

```
MICRO_REGISTRY=mdns micro call digitalocean Droplets.Create '{"name":"Foo","image":"ubuntu-14-04-x64","region":"nyc3","size":"s-1vcpu-1gb"}'
```

3. Get

```
MICRO_REGISTRY=mdns micro call digitalocean Droplets.Get '{"id":118550799}'
```

4. Delete

```
MICRO_REGISTRY=mdns micro call digitalocean Droplets.Get '{"id":118550799}'
```

# TODO

1. Decide if `proto` directory is ok or not?

    https://github.com/golang-standards/project-layout/ has no `proto` directory.
    I assume we could copy Helm's layout which uses `_proto` for `*.proto` files and `pkg/proto` for compiled client code.
    I'm open for discussion. For now, leaving `proto` and `pkg/proto` as it's easier to script it this way.

2. Decide if we should rename `proto` and `service` packages to be under the same main `package`

    Currently we have:

        ```
        proto "github.com/jerryhoio/digitalocean/pkg/proto/greeter"
      	svc "github.com/jerryhoio/digitalocean/pkg/service/greeter"
        ```

    Maybe a better structure would be:

        ```
        proto "github.com/jerryhoio/digitalocean/pkg/greeter/proto"
      	svc "github.com/jerryhoio/digitalocean/pkg/greeter/service"
        ```

3. Logging - decide if and how to log events.
