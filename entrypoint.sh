#!/bin/bash

if [ "$VERBOSE" = "true" ]; then
  nohup /etc/caddy/status > /etc/caddy/status.log 2>&1 &

  sleep 3
  cat /etc/caddy/status.log
else
  nohup /etc/caddy/status > /dev/null 2>&1 &
fi

caddy file-server --root /etc/caddy --listen :8080