package models

import "github.com/globalsign/mgo/bson"

type Ticket struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	ProductID bson.ObjectId `bson:"product_id" json:"product_id"`
	Price int `bson:"price" json:"price"`
	Status string `bson:"status" json:"status"`
}

