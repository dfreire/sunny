#!/bin/bash
source env.production.sh
URL="http://localhost$SUNNY_PORT/send-to-newsletter?appToken=$SUNNY_APP_TOKEN"
curl -XPOST $URL &> /dev/null
