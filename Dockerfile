FROM golang:1.22.5-alpine3.20 AS build-dev

WORKDIR /go/src/app

COPY --link service/go.sum service/go.mod ./

RUN apk add --no-cache upx || \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy

COPY --link service/main.go service/go.mod service/go.sum ./

RUN go build -o status . && \
    [ -e /usr/bin/upx ] && upx status || echo

FROM caddy:2.4.5-alpine

WORKDIR /etc/caddy

COPY --link src ./src
COPY --link Caddyfile config.cfg favicon.ico index.html entrypoint.sh ./
COPY --from=build-dev /go/src/app/status ./status

RUN apk update && \
    apk add --no-cache bash curl && \
    chmod +x /etc/caddy/status && \
    chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]