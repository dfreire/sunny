#!/bin/bash
source env.production.sh
HERE=`cd -P $(dirname $0) && pwd`
mkdir -p $HERE/log
nohup $HERE/sunny &> $HERE/log/log.txt < /dev/null &
