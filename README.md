# go-gin-starter


## Commands

Running `make` will display a list of available commands:

```bash
$ make
```

Output:

```bash
usage: make (sub-commands ...)

Gin Starter commands.

commands:
help                           Display help
vendor                         Vendor the application
docker                         Build the docker image
start                          Start the application
inspect                        View the docker image tags
test                           Run unit test
```

## Installation

```bash
$ go get -u golang.org/x/vgo
```

## vgo

**vgo** is used for dependency management, as well as building/testing the application.

```bash
# Creates a go.mod file locally
$ vgo mod -init

# Vendor package
$ vgo mod -vendor

# Sync packages
$ vgo mod -sync

# Update packages
$ vgo get -u

# Build using packages from vendor folder
$ vgo build -getmode=vendor -o app
```

## Run

```bash
$ go run main.go
```

## Docker


```bash
# Build docker image
$ make docker

# View docker-labels of the built image
$ make inspect | jq
```

Output:

```json
{
  "org.label-schema.build-date": "2018-07-17T14:13:52Z",
  "org.label-schema.description": "go gin starter application",
  "org.label-schema.docker.cmd": "docker run -d alextanhongpin/go-gin-starter",
  "org.label-schema.docker.schema-version": "1.0",
  "org.label-schema.name": "go-gin-starter",
  "org.label-schema.url": "https://example.com",
  "org.label-schema.vcs-ref": "bdeae5be657a54790909b99a3d1eddd719b7830a",
  "org.label-schema.vcs-url": "https://github.com/alextanhongpin/go-gin-starter.git",
  "org.label-schema.vendor": "alextanhongpin",
  "org.label-schema.version": "bdeae5b"
}
```

To view the docker image size:

```bash
$ docker images | grep go-gin-starter
```

Output:

```bash
alextanhongpin/go-gin-starter      latest              0dbd3a78cef6        23 seconds ago      17.8MB
```

## TODO

- event logging (for webhook)
- logger
- cors
- secure
- cancellation/retry/circuit breaker
- tracing