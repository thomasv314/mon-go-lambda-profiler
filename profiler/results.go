package profiler

import (
	"encoding/json"
	"github.com/thomasv314/mongo-tools/common/bsonutil"
	"log"
)

func (p Profiler) ProcessResults() (err error) {
	results, err := p.results()
	if err != nil {
		log.Println("error querying results", err)
		return
	}

	results, err = p.filterResults(results)
	if err != nil {
		log.Println("error filtering results", err)
		return
	}

	filteredResultsJson, err := p.resultsAsJSON(results)
	if err != nil {
		log.Println("error marshalling to json", err)
		return
	}

	err = p.uploadToS3(filteredResultsJson)
	if err != nil {
		log.Println("error uploading to s3", err)
		return
	}
	return
}

func (p Profiler) results() (result []interface{}, err error) {
	result = make([]interface{}, 0)
	err = p.Db.C("system.profile").Find(nil).All(&result)
	return
}

func (p Profiler) resultsAsJSON(results []interface{}) (jsonBytes []byte, err error) {
	jsonArr := make([]interface{}, len(results))

	for r := range results {
		asJson, err := bsonutil.GetBSONValueAsJSON(results[r])
		if err != nil {
			panic(err) // unsure what to do yet
		}

		jsonArr[r] = asJson
	}

	jsonBytes, err = json.Marshal(jsonArr)
	return
}
