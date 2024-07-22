#!/bin/bash

nohup /etc/caddy/status > /dev/ull 2>&1 &

caddy file-server --root /etc/caddy