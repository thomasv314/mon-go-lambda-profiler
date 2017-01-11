package profiler

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)

func (p Profiler) filterResults(results []interface{}) (filteredResults []interface{}, err error) {
	filteredResults = make([]interface{}, len(results))

	for i := range results {
		bsonResult := results[i].(bson.M)

		if bsonResult["query"] != nil {
			bsonQuery := bsonResult["query"].(bson.M)
			bsonResult["query"] = filterValuesFromBSON(bsonQuery)
		}

		log.Println(i, bsonResult)

		filteredResults[i] = bsonResult
	}

	return
}

func filterValuesFromBSON(doc bson.M) bson.M {
	if doc == nil {
		return doc
	}

	for key, val := range doc {
		if val != nil {
			valType := reflect.TypeOf(doc[key]).String()
			switch valType {
			case "bson.M":
				doc[key] = filterValuesFromBSON(val.(bson.M))
			default:
				doc[key] = nil
			}
		}
	}
	return doc
}
