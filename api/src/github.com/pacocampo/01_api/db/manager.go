package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

//GetSession
//returns a mongodb session for current request
func GetSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal("Database connection error")
		panic(err)
	}
	return session
}
