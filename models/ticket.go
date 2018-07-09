package models

import "github.com/globalsign/mgo/bson"

type Ticket struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID bson.ObjectId `bson:"product_id,omitempty" json:"product_id,omitempty"`
	Price int `bson:"price,omitempty" json:"price,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
}

