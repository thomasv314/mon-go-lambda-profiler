#!/bin/bash

aws lambda update-function-code \
  --function-name $LAMBDA_NAME \
  --zip-file fileb://handler.zip
