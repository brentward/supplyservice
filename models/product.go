package models

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	BasePrice int64 `bson:"base_price,omitempty" json:"base_price,omitempty"`
	ExternalPartnerId string `bson:"external_partner_id,omitempty" json:"external_partner_id,omitempty"`
	Type string `bson:"type,omitempty" json:"type,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty"`
}
