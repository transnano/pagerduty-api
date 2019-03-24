# pagerduty-api

``` sh
$ go get -u github.com/gin-gonic/gin
```

## Go version

``` shell
$ go version
go version go1.11.4 darwin/amd64
```

## Init

``` shell
$ GO111MODULE=on go mod init
go: creating new go.mod: module gitlab.com/transnano/pagerduty-api
```

## Build

``` shell
$ GO111MODULE=on go build main.go
$ GO111MODULE=on go run main.go
# exec sample
$ curl -X POST -H "Content-Type: application/json" -d '{"aaa": "test"}' -w '%{http_code}\n' http://localhost:8080/v2/enqueue
{"dedup_key":"samplekeyhere","message":"Event processed","status":"success"}202

$ curl -X GET http://localhost:8080/metrics
...
go_info{version="go1.11.4"} 1
...
```
