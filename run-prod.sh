#!/bin/bash

HERE=`cd -P $(dirname $0) && pwd`

export SUNNY_ENV="production"
export SUNNY_JWT_KEY="2bf92a58-b8e3-48c4-8afd-483d4618d13d"
export SUNNY_SQLITE_DB="$HERE/sunny.db"
export SUNNY_PORT=":3500"

./sunny
