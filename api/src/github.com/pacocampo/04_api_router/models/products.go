package models

import "gopkg.in/mgo.v2/bson"

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float32       `json:"price" bson:"price"`
}
