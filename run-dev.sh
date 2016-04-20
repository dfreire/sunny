#!/bin/bash

HERE=`cd -P $(dirname $0) && pwd`

#export LOGXI=*

export SUNNY_ENV="development"
export SUNNY_JWT_KEY="451b047b-9cc6-4c76-9c5c-63e80b23e7d6"
export SUNNY_SQLITE_DB="$HERE/sqlite.db"
export SUNNY_PORT=":3500"

# rm -rf $SUNNY_SQLITE_DB
go run main.go
