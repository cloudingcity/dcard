# Dcard

[![Build Status](https://travis-ci.com/cloudingcity/dcard.svg?branch=master)](https://travis-ci.com/cloudingcity/dcard)
[![codecov](https://codecov.io/gh/cloudingcity/dcard/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/dcard)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/dcard)](https://goreportcard.com/report/github.com/cloudingcity/dcard)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/cloudingcity/dcard)

## Usage

```shell script
curl -i https://dcard-shawnpeng.herokuapp.com
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
