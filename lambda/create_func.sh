#!/bin/bash

aws lambda create-function \
  --role arn:aws:iam::$AWS_ACCOUNT_ID:role/lambda_basic_execution \
  --function-name $LAMBDA_NAME \
  --zip-file fileb://handler.zip \
  --runtime python2.7 \
  --handler handler.Handle
