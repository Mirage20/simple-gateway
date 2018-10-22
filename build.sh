#!/usr/bin/env bash

export CGO_ENABLED=0

go build -o simple-gateway -x ./cmd/gateway

docker build -t mirage20/simple-gateway .

docker push mirage20/simple-gateway
