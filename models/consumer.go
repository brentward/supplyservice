package models

import "github.com/globalsign/mgo/bson"

type Consumer struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty"`
	Role string `bson:"role,omitempty" json:"role,omitempty"`
	PartnerIDs []bson.ObjectId `bson:"partners,omitempty" json:"partners",omitempty`
}
