FROM golang:1.22.5-alpine3.20 AS build-dev

WORKDIR /go/src/app

COPY --link service/go.sum service/go.mod ./

RUN apk add --no-cache upx || \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy

COPY --link service/main.go service/go.mod service/go.sum ./
COPY --link service/notify ./notify

RUN go build -o status . && \
    [ -e /usr/bin/upx ] && upx status || echo

FROM caddy:2.8.4-alpine

WORKDIR /etc/caddy

COPY --link logs ./logs
COPY --link public ./public
COPY --link src ./src
COPY --link Caddyfile config.cfg favicon.ico index.html entrypoint.sh env.yaml password.Caddyfile ./
COPY --from=build-dev /go/src/app/status ./status

RUN apk update && \
    apk add --no-cache bash tzdata && \
    chmod +x /etc/caddy/status && \
    chmod +x ./entrypoint.sh && \
    chmod 777 /etc/caddy/logs

ENV VERBOSE=false
ENV TZ=Asia/Shanghai

ENTRYPOINT ["./entrypoint.sh"]