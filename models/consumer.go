package models

import "github.com/globalsign/mgo/bson"

type Consumer struct {
	ID bson.ObjectId `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Active bool `json:"active"`
	Role string `json:"role"`
	PartnerID []bson.ObjectId `json:"partner_id"`
}
