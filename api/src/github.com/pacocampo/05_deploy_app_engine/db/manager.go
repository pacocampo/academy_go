package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host = "mongodb://paco:41000705@ds127536.mlab.com:27536/db-api"
	db   = "db-api"
)

// struct Database struct {
// 	session *mgo.Session
// }

// func NewSession() *Database {
// 	s := GetSession()
// 	return Database(s)
// }
//GetSession
//returns a mongodb session for current request
func GetSession() *mgo.Session {

	session, err := mgo.Dial(host)
	if err != nil {
		log.Fatal("Database connection error", err)
		panic(err)
	}
	return session
}

func GetAll(session *mgo.Session, collection string, object interface{}) error {
	return session.DB(db).C(collection).Find(nil).All(object)
}

func Get(session *mgo.Session, collection, id string, object interface{}) error {
	oid := bson.ObjectIdHex(id)
	return session.DB(db).C(collection).FindId(oid).One(object)
}

func Create(session *mgo.Session, collection string, object interface{}) error {
	return session.DB(db).C(collection).Insert(object)
}

func Update(session *mgo.Session, collection string, id bson.ObjectId, object interface{}) error {
	return session.DB(db).C(collection).UpdateId(id, object)
}
