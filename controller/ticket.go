package controller

import (
	. "supplyservice/models"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
)

func GetTicket(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	ticket, err := dao.FindTicketById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, ticket)
}

func GetTickets(w http.ResponseWriter, _ *http.Request)  {
	tickets, err := dao.FindAllTickets()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, tickets)
}

func CreateTicket(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var ticket Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	ticket.ID = bson.NewObjectId()
	if err := dao.InsertTicket(ticket); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, ticket)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	defer r.Body.Close()
	var ticket Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateTicket(params["id"], ticket); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeleteTicket(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	if err := dao.DeleteTicket(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}
