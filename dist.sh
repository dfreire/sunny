#!/bin/bash

docker-machine start sunny
eval "$(docker-machine env sunny)"

HERE=`cd -P $(dirname $0) && pwd`

docker build --tag sunny .
docker run -it --rm -v $HERE:/go/src/github.com/dfreire/sunny sunny go build -o ./dist/sunny main.go

rm -rf $HERE/dist/templates
mkdir -p $HERE/dist/templates
cp -R $HERE/templates $HERE/dist/
