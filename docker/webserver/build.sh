#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t danielmerchant/testserver .
docker push danielmerchant/testserver