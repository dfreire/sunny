#!/bin/bash
docker run -it --rm -v "$PWD":/go/src/github.com/dfreire/sunny sunny go build -o ./dist/sunny main.go
