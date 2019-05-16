#!/bin/bash

set -e

rm -f main.zip go-imports
GOOS=linux GOARCH=amd64 GO111MODULE=on go build
zip -X -r ./main.zip go-imports

aws lambda update-function-code --function-name dedis-go-imports --zip-file fileb://main.zip
