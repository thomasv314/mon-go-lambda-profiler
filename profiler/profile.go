package profiler

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Profiler struct {
	Session        *mgo.Session
	Db             *mgo.Database
	MaxQueryTimeMs int
}

type Result struct{}

func Create(dbUrl, dbName string, maxQueryTimeMs int) Profiler {
	log.Println(dbUrl, dbName, maxQueryTimeMs)

	session, err := mgo.Dial(dbUrl)

	if err != nil {
		log.Println("Session open Error:", err)
		panic(err)
	}

	db := session.DB(dbName)

	return Profiler{
		Session:        session,
		Db:             db,
		MaxQueryTimeMs: maxQueryTimeMs,
	}
}

func (p Profiler) Results() (result []interface{}, err error) {
	result = make([]interface{}, 0)
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
	log.Println("Closed Mongo Connection.")
}
