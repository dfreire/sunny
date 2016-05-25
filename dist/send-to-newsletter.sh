#!/bin/bash
HERE=`cd -P $(dirname $0) && pwd`

APP_PORT=`cat config.production.yaml | grep "^port:" | awk '{ print $2 }'`
APP_TOKEN=`cat config.production.yaml | grep "^appToken:" | awk '{ print $2 }'`

URL="http://localhost$APP_PORT/send-to-newsletter?appToken=$APP_TOKEN"

curl -XPOST $URL &> /dev/null
 