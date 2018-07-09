package controller

import (
	. "supplyservice/models"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
)

func GetTransaction(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	transaction, err := dao.FindTransactionById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, transaction)
}

func GetTransactions(w http.ResponseWriter, _ *http.Request)  {
	transactions, err := dao.FindAllTransactions()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	transaction.ID = bson.NewObjectId()
	if err := dao.InsertTransaction(transaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, transaction)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	defer r.Body.Close()
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateTransaction(params["id"], transaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeleteTransaction(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	if err := dao.DeleteTransaction(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}
