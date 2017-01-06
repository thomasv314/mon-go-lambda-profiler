# mon-go-lambda-profiler

Use an AWS lambda function to profile MongoDB queries.

Lambda execution timeout has to be longer than the mongo profile duration seconds, with a
little bit of a buffer to allow uploading the profile result JSON to S3.

## Environment Variables

```
export AWS_ACCOUNT_ID="myaccid"                  # account id to upload func to
export LAMBDA_NAME="mongo-lambda-profiler"       # the name of the lambda func
export MONGO_PROFILE_URL="localhost:27017"       # mongo url
export MONGO_PROFILE_DB_NAME="from_development"  # mongo db
export MONGO_PROFILE_MAX_QUERY_MS="1"            # min ms to consider a query slow
export MONGO_PROFILE_DURATION_SECS="5"           # how long to run your sample
```

## Example Output

```
START RequestId: 38ab12dd-d426-11e6-a62a-b145c1ad378a Version: $LATEST
2017-01-06T15:38:50.914Z	Profiling mysite_production at 10.37.18.98:27000 for queries slower than 5ms
2017-01-06T15:38:50.914Z	10.37.18.98:27000 mysite_production 5
2017-01-06T15:38:51.014Z	Profiling enabled.. Waiting 10 seconds.
2017-01-06T15:39:01.697Z	Uploaded to S3
2017-01-06T15:39:01.697Z	Closed Mongo Connection.
END RequestId: 38ab12dd-d426-11e6-a62a-b145c1ad378a
REPORT RequestId: 38ab12dd-d426-11e6-a62a-b145c1ad378a	Duration: 10793.75 ms	Billed Duration: 10800 ms 	Memory Size: 128 MB	Max Memory Used: 22 MB	
```
