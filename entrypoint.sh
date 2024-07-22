#!/bin/bash

if [ "$VERBOSE" = "false" ]; then
  nohup /etc/caddy/status > /dev/ull 2>&1 &
else
  nohup /etc/caddy/status > /etc/caddy/status.log 2>&1 &
fi

caddy file-server --root /etc/caddy --listen :8080