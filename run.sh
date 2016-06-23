#!/bin/bash
source env.development.sh
mkdir -p $SUNNY_TMP_DIR
go run main.go
