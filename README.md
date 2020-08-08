# Rate Limit HTTP Server

[![Build Status](https://travis-ci.com/cloudingcity/ratelimit-server.svg?branch=master)](https://travis-ci.com/cloudingcity/ratelimit-server)
[![codecov](https://codecov.io/gh/cloudingcity/ratelimit-server/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/ratelimit-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/ratelimit-server)](https://goreportcard.com/report/github.com/cloudingcity/ratelimit-server)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/cloudingcity/ratelimit-server)

A simple rate limit HTTP server limit every client IP

## Usage

```shell script
curl -i https://ratelimit-server-shawnpeng.herokuapp.com
```

### Environment

ENV | DESCRIPTION
--- | ---
PORT | http server listen port (default: `8080`)
RATE_LIMIT | Max number of requests furing `RATE_TIMEOUT` seconds (default: `60`)
RATE_TIMEOUT | How long keep requests in seconds (default: `60`)

## Development

```shell script
go run main.go
```

### Test

```shell script
make test
```

### Lint

```shell script
make lint
```
