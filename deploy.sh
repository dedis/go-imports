#!/bin/bash

set -e

rm -f main.zip go-imports
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=auto go build
zip -X -r ./main.zip go-imports

aws lambda update-function-code --function-name dedis-go-imports --zip-file fileb://main.zip
