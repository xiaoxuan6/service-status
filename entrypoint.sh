#!/bin/bash

if [ "$VERBOSE" = "true" ]; then
  nohup /etc/caddy/status > /tmp/status.log 2>&1 &

  sleep 3
  cat /tmp/status.log
else
  nohup /etc/caddy/status > /dev/null 2>&1 &
fi

if [[ -n "$USERNAME" && -n "$PASSWORD" ]]; then
  NEW_PASSWORD=$(caddy hash-password -p "$PASSWORD")
  export USERNAME=$USERNAME
  export PASSWORD=$NEW_PASSWORD

  caddy run --config=/etc/caddy/password.Caddyfile
else
  caddy run --config=/etc/caddy/Caddyfile
fi