# pagerduty-api

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
go: creating new go.mod: module gitlab.com/transnano/pagerduty-api
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
