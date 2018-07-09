package models

import "github.com/globalsign/mgo/bson"

type Partner struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty"`
	ProductIDs []bson.ObjectId `bson:"products,omitempty" json:"products,omitempty"`
}