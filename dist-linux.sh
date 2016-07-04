#!/bin/bash

go build -o ./dist/sunny main.go

export HERE=`cd -P $(dirname $0) && pwd`
rm -rf $HERE/dist/templates
mkdir -p $HERE/dist/templates
cp -R $HERE/templates $HERE/dist/
