package models

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID bson.ObjectId `json:"id"`
	TicketID bson.ObjectId `json:"ticket_id"`
	Status string `json:"status"`
}
