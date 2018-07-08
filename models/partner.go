package models

import "github.com/globalsign/mgo/bson"

type Partner struct {
	ID bson.ObjectId `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Active bool `json:"active"`
	TicketID []bson.ObjectId `json:"tickets"`
}