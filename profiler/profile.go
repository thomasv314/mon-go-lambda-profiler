package profiler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Profiler struct {
	Session        *mgo.Session
	Db             *mgo.Database
	MaxQueryTimeMs int
}

type Result struct{}

func Create(dbUrl, dbName string, maxQueryTimeMs int) Profiler {
	fmt.Println(dbUrl, dbName, maxQueryTimeMs)

	session, err := mgo.Dial(dbUrl)

	if err != nil {
		fmt.Println("Session open Error:", err)
		panic(err)
	}

	db := session.DB(dbName)

	return Profiler{
		Session:        session,
		Db:             db,
		MaxQueryTimeMs: maxQueryTimeMs,
	}
}

func (p Profiler) Results() (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 0)
	err = p.Db.C("system.profile").Find(nil).All(&result)
	return
}

func (p Profiler) DisableProfiling() (result Result, err error) {
	err = p.Db.Run(bson.D{{"profile", 0}}, &result)
	return
}

func (p Profiler) EnableProfiling() (result Result, err error) {
	err = p.Db.Run(bson.D{{"profile", 2}, {"slowms", p.MaxQueryTimeMs}}, &result)
	return
}

func (p Profiler) Close() {
	p.Session.Close()
	fmt.Println("Closed Mongo Connection.")
}
