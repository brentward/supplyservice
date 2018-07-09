package models

import "github.com/globalsign/mgo/bson"

type Transaction struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	ConsumerID bson.ObjectId `bson:"consumer_id,omitempty" json:"consumer_id,omitempty"`
	Consumer Consumer
	PartnerID bson.ObjectId `bson:"partner_id,omitempty" json:"partner_id,omitempty"`
	TicketIDs []bson.ObjectId `bson:"tickets,omitempty" json:"tickets,omitempty"`
}
