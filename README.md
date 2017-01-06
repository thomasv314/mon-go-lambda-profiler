# mon-go-lambda-profiler

Use an AWS lambda function to profile MongoDB queries

```
export MONGO_PROFILE_URL="localhost:27017"       # mongo url
export MONGO_PROFILE_DB_NAME="from_development"  # mongo db
export MONGO_PROFILE_MAX_QUERY_MS="1"            # min ms to consider a query slow
export MONGO_PROFILE_DURATION_SECS="5"           # how long to run your sample
```
