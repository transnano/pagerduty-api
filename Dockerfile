FROM golang:1.15.0
WORKDIR /go/src/github.com/transnano/pagerduty-api/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o pagerduty-api -ldflags "-s -w \
-X 'main.Version=$(git describe --tags --abbrev=0)'"

FROM alpine:3.12.0
LABEL maintainer="Transnano <transnano.jp@gmail.com>"
RUN apk --no-cache add ca-certificates
EXPOSE 8080
COPY --from=0 /go/src/github.com/transnano/pagerduty-api/pagerduty-api /bin/pagerduty-api
ENTRYPOINT ["/bin/pagerduty-api"]
