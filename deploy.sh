#!/bin/bash

rm main.zip go-imports
GO111MODULE=on go build
zip -X -r ./main.zip go-imports

aws lambda update-function-code --function-name dedis-go-imports --zip-file fileb://main.zip
