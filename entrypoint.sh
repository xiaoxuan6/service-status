#!/bin/bash

nohup /etc/caddy/status > /dev/null 2>&1 &

caddy file-server --root /etc/caddy --listen :8080