package models

import "github.com/globalsign/mgo/bson"

type Transaction struct {
	ID bson.ObjectId `json:"id"`
	Status string `json:"name"`
	ConsumerID bson.ObjectId `json:"consumer_id"`
	CatalogID bson.ObjectId `json:"catalog_id"`
	ProductID []bson.ObjectId `json:"products"`
}
