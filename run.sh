#!/bin/bash
source env.development.sh
mkdir -p $SUNNY_MAILER_ATTACHMENTS_DIR
go run main.go
