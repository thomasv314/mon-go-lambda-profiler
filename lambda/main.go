package main

// /* Required, but no C code needed. */
import "C"

import (
	"encoding/json"
	"errors"
	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	"github.com/thomasv314/mon-go-lambda-profiler/profiler"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_MAX_QUERY_MS          int = 50
	DEFAULT_PROFILE_DURATION_SECS int = 30
)

var (
	url                 string
	dbName              string
	maxQueryTimeMs      int
	profileDurationSecs int
)

func init() {
	runtime.HandleFunc(handle)
}

func main() {}

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	setProfileConfig()
	err := setRequiredEnvVars()
	if err != nil {
		log.Println("Missing required EnvVars.")
		return "envvars missing", err
	}

	log.Println("Profiling", dbName, "at", url, "for queries slower than", maxQueryTimeMs)

	prof := profiler.Create(url, dbName, maxQueryTimeMs)
	defer prof.Close()

	_, err = prof.EnableProfiling()

	if err != nil {
		return "error", err
	} else {
		log.Println("Profiling enabled.. Waiting", profileDurationSecs, "seconds.")
	}

	select {
	case <-time.After(time.Duration(profileDurationSecs) * time.Second):
		_, err = prof.DisableProfiling()
		if err != nil {
			log.Println("Profiling could not be disabled.. Manual intervention may be wise.")
			panic(err)
		}

		prof.UploadResultsToS3()
	}

	return "it worked?", nil
}

func setProfileConfig() {
	maxQueryTimeMsStr, err := checkForEnvVar("MONGO_PROFILE_MAX_QUERY_MS")
	maxQueryTimeMs, err = strconv.Atoi(maxQueryTimeMsStr)
	if err != nil {
		maxQueryTimeMs = DEFAULT_MAX_QUERY_MS
	}

	profileDurationSecsStr, err := checkForEnvVar("MONGO_PROFILE_DURATION_SECS")
	profileDurationSecs, err = strconv.Atoi(profileDurationSecsStr)
	if err != nil {
		profileDurationSecs = DEFAULT_PROFILE_DURATION_SECS
	}
}

func setRequiredEnvVars() (err error) {
	url, err = checkForEnvVar("MONGO_PROFILE_URL")
	if err != nil {
		return err
	}

	dbName, err = checkForEnvVar("MONGO_PROFILE_DB_NAME")
	if err != nil {
		return err
	}

	return nil
}

func checkForEnvVar(name string) (value string, err error) {
	value = os.Getenv(name)
	if value == "" {
		err = errors.New("Missing environment variable: " + name)
	}
	return
}
