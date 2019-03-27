# Golang Redis Rate Limiter [![Go Report Card](https://goreportcard.com/badge/github.com/shredx/golang-api-gateway)](https://goreportcard.com/report/github.com/shredx/golang-api-gateway)

It's API limiter built in Go while we were trying out Redis.
It has got dependency on a node app to generate tokens [node-redis-rate-limiter](https://github.com/shredx/node-redis-rate-limiter)
and [](https://github.com/shredx/golang-redis-rate-limiter).


## Getting Started

### Prerequisite
#### Local Setup
* [Go](https://golang.org/doc/install) -- Development environment
* [dep](https://golang.github.io/dep/docs/installation.html) -- Dependency management
* [Redis](https://redis.io/download) -- Cache storage

#### Docker
* [Docker](https://www.docker.com/products/docker-desktop)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Installation
#### Local Setup
```sh
go get -u github.com/shredx/golang-api-gateway
cd $GOPATH/github.com/shredx/golang-api-gateway
dep ensure
go run main.go
```
#### Docker
```sh
git clone https://github.com/shredx/golang-api-gateway
cd golang-api-gateway
git submodule init
git submodule update
docker-compose up
```

#### Help
```sh
go run main.go -help
```
### Usage
* Refer the [usage doc](./Usage.md)

## Architecture
The architecture might not be perfect as it was just a weekend project to explore Redis.

![Architecture](https://github.com/shredx/golang-redis-rate-limiter/raw/master/.github/architecture.png)
