FROM golang:1.18-buster as builder

RUN set -ex && \
    apt-get -qq update && apt-get -y -qq install \
      gcc

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -a -v -o app cmd/main.go

# main
FROM alpine:latest

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["app"]

WORKDIR /app

COPY --from=builder /go/src/app/app /usr/local/bin/
COPY config/*.yaml /app/config/