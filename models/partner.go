package models

import "github.com/globalsign/mgo/bson"

type Partner struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Active bool `bson:"active" json:"active"`
	ProductIDs []bson.ObjectId `bson:"products" json:"products"`
}