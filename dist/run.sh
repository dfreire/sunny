#!/bin/bash
HERE=`cd -P $(dirname $0) && pwd`
mkdir -p $HERE/db
mkdir -p $HERE/log
export ENV="production"
nohup $HERE/sunny &> $HERE/log/log.txt < /dev/null &
