#!/bin/bash
export ENV="development"
go test -v -coverprofile=coverage.out ./handlers/...
go tool cover -func=coverage.out
