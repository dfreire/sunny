#!/bin/bash
source env.production.sh
mkdir -p $SUNNY_MAILER_ATTACHMENTS_DIR
mkdir -p $HERE/log
nohup $HERE/sunny &> $HERE/log/log.txt < /dev/null &
