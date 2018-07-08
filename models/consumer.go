package models

import "github.com/globalsign/mgo/bson"

type Consumer struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Active bool `bson:"active" json:"active"`
	Role string `bson:"role" json:"role"`
	PartnerIDs []bson.ObjectId `bson:"partners" json:"partners"`
}
