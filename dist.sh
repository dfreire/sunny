#!/bin/bash

HERE=`cd -P $(dirname $0) && pwd`

mkdir -p $HERE/dist

docker build --tag sunny .
docker run -it --rm -v $HERE:/go/src/github.com/dfreire/sunny sunny go build -o ./dist/sunny main.go
