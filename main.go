package main

import (
	"github.com/thomasv314/mon-go-lambda-profiler/profiler"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	url := os.Getenv("MONGO_PROFILE_URL")
	dbName := os.Getenv("MONGO_PROFILE_DB_NAME")
	maxQueryTimeMsStr := os.Getenv("MONGO_PROFILE_MAX_QUERY_MS")
	profileDurationSecsStr := os.Getenv("MONGO_PROFILE_DURATION_SECS")

	maxQueryTimeMs, err := strconv.Atoi(maxQueryTimeMsStr)
	handleErr(err)

	profileDurationSecs, err := strconv.Atoi(profileDurationSecsStr)
	handleErr(err)

	prof := profiler.Create(url, dbName, maxQueryTimeMs)
	defer prof.Close()

	_, err = prof.EnableProfiling()

	if err != nil {
		handleErr(err)
	} else {
		log.Println("Profiling enabled.. Waiting", profileDurationSecsStr, "seconds.")
	}

	select {
	case <-time.After(time.Duration(profileDurationSecs) * time.Second):
		_, err = prof.DisableProfiling()
		if err != nil {
			log.Println("Profiling could not be disabled.. Manual intervention may be wise.")
			panic(err)
		}

		log.Println("Profiling disabled.")

		err := prof.ProcessResults()

		if err != nil {
			log.Println("Error processing results:", err)
			panic(err)
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
