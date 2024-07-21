#!/bin/bash

nohup /etc/caddy/status > /etc/caddy/status.log 2>&1 &

caddy file-server --root /etc/caddy