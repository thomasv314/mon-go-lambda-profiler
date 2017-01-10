package profiler

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)

func (p Profiler) filterResults(results []interface{}) (filteredResults []interface{}, err error) {
	filteredResults = make([]interface{}, len(results))

	filterSingleResult(results[0])
	/*	for i := range results {
			if reflect.TypeOf(results[i]).String() != "bson.M" {
				panic("attempted to filter result that was not of bson.M type")
			}
			filteredResults[i] = filterSingleResult(results[i])
		}
	*/
	return
}

func filterSingleResult(result interface{}) interface{} {
	bsonResult := result.(bson.M)

	for key, val := range bsonResult {
		if val != nil {
			valueType := reflect.TypeOf(val)

			if valueType == nil {
				bsonResult[key] = nil
			} else if valueType.String() == "bson.M" {
				bsonResult[key] = filterSingleResult(val)
			} else {
				bsonResult[key] = reflect.Zero(valueType)
			}
			log.Println("["+reflect.TypeOf(val).String()+"] "+key, ":", val)
		} else {
			log.Println("[nil] "+key, ":", val)
			bsonResult[key] = ""
		}
	}

	return bsonResult
}
