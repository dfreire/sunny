#!/bin/bash

HERE=`cd -P $(dirname $0) && pwd`

mkdir -p $HERE/dist
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $HERE/dist/sunny main.go
