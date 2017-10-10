#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t danielmerchant/zipsvr .
docker push danielmerchant/zipsvr