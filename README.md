# go-gin-starter


## Installation

```bash
$ go get -u golang.org/x/vgo
```

## vgo

```bash
# Vendor package
$ vgo mod -vendor

# Sync packages
$ vgo mod -sync

# Update packages
$ vgo get -u
```

## Run

```bash
$ go run main.go
```

## TODO

- event logging (for webhook)
- logger
- cors
- secure
- cancellation/retry/circuit breaker
- tracing