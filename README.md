[![CircleCI](https://circleci.com/gh/alexfalkowski/go-monolith.svg?style=svg)](https://circleci.com/gh/alexfalkowski/go-monolith)
[![codecov](https://codecov.io/gh/alexfalkowski/go-monolith/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/go-monolith)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/go-monolith)](https://goreportcard.com/report/github.com/alexfalkowski/go-monolith)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/go-monolith.svg)](https://pkg.go.dev/github.com/alexfalkowski/go-monolith)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Monolith

There has been tremendous amounts of time being spent on creating [microservices](https://microservices.io/), though is this the place to start?

It is not a trick question, of course not. It is best to start with a monolith and transition to [microservices](https://martinfowler.com/articles/break-monolith-into-microservices.html) when needed.

## Background

After many years of seeing [distributed monoliths](https://www.gremlin.com/blog/is-your-microservice-a-distributed-monolith). This could potentially be a compromise.

### Why a service?

This project allows us to use a [mono repository](https://monorepo.tools/) approach to build connected services.

## Server

This no different to the other services we are built, they just contain it all together.

### API

The [api](api) are where we define out [protobuf](https://protobuf.dev/) services.

### Servers

All the [servers](internal/api) are defined in a [usual way](https://grpc.io/docs/languages/go/basics/).

Each of the services talks to each other via a [client](internal/api/client). As all of this is running through [localhost](https://en.wikipedia.org/wiki/Localhost).

Each service is reachable by defining [REST](https://github.com/alexfalkowski/go-service/tree/master/net/http/rest) endpoints.

## Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

## Other Systems

[gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) works in a similar way.

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:

- [Ruby](https://www.ruby-lang.org/en/)
- [Golang](https://go.dev/)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
