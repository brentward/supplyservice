package models

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	BasePrice int64 `bson:"base_price" json:"base_price"`
	ExternalPartnerId string `bson:"external_partner_id" json:"external_partner_id"`
	Type string `bson:"type" json:"type"`
	Active bool `bson:"active" json:"active"`
}
