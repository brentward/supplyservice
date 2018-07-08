package models

import "github.com/globalsign/mgo/bson"

type Transaction struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Status string `bson:"name" json:"name"`
	ConsumerID bson.ObjectId `bson:"consumer_id" json:"consumer_id"`
	PartnerID bson.ObjectId `bson:"partner_id" json:"partner_id"`
	TicketIDs []bson.ObjectId `bson:"tickets" json:"tickets"`
}
