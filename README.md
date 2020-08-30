# pagerduty-api ![Releases](https://github.com/transnano/pagerduty-api/workflows/Releases/badge.svg) ![Publish Docker image](https://github.com/transnano/pagerduty-api/workflows/Publish%20Docker%20image/badge.svg) ![Vulnerability Scan](https://github.com/transnano/pagerduty-api/workflows/Vulnerability%20Scan/badge.svg)

![License](https://img.shields.io/github/license/transnano/pagerduty-api?style=flat)

![Container image version](https://img.shields.io/docker/v/transnano/pagerduty-api/latest?style=flat)
![Container image size](https://img.shields.io/docker/image-size/transnano/pagerduty-api/latest?style=flat)
![Container image pulls](https://img.shields.io/docker/pulls/transnano/pagerduty-api?style=flat)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/transnano/pagerduty-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/transnano/pagerduty-api)](https://goreportcard.com/report/github.com/transnano/pagerduty-api)

``` sh
$ go get -u github.com/gin-gonic/gin
```

## Go version

``` shell
$ go version
go version go1.15 linux/amd64
```

## Init

``` shell
$ go mod init
go: creating new go.mod: module github.com/transnano/pagerduty-api
```

## Build

``` shell
$ go build main.go
$ go run main.go
# exec sample
$ curl -X POST -H "Content-Type: application/json" -d '{"aaa": "test"}' -w '%{http_code}\n' http://localhost:8080/v2/enqueue
{"dedup_key":"samplekeyhere","message":"Event processed","status":"success"}202

$ curl -X GET http://localhost:8080/metrics
...
go_info{version="go1.15"} 1
...
```
