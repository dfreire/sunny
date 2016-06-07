#!/bin/bash
HERE=`cd -P $(dirname $0) && pwd`

#export SUNNY_ENV="test"
#export SUNNY_DEBUG="true"
#export SUNNY_APP_TOKEN="2fe9a70a-46f2-4d00-88f2-6f66ed903426"
#export SUNNY_DATABASE="development.db"
#export SUNNY_PORT=":3500"
#export SUNNY_MAILER="log"
export SUNNY_MAILER_TEMPLATES_DIR="$HERE/templates"
export SUNNY_TEAM_EMAIL="team-6f66ed903426@mailinator.com"
export SUNNY_OWNER_EMAIL="owner-6f66ed903426@mailinator.com"
export SUNNY_NOTIFICATION_EMAILS="a-6f66ed903426@mailinator.com,b-6f66ed903426@mailinator.com"

go test -v -coverprofile=coverage.out ./commands/...
# go test -v -coverprofile=coverage.out ./handlers/...
go tool cover -func=coverage.out
