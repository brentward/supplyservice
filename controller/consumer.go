package controller

import (
	. "supplyservice/models"
	"net/http"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func GetConsumer(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	consumer, err := dao.FindConsumerById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, consumer)
}

func GetConsumers(w http.ResponseWriter, _ *http.Request)  {
	consumers, err := dao.FindAllConsumers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, consumers)
}

func CreateConsumer(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var consumer Consumer
	if err := json.NewDecoder(r.Body).Decode(&consumer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	consumer.ID = bson.NewObjectId()
	if err := dao.InsertConsumer(consumer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, consumer)
}

func UpdateConsumer(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var consumer Consumer
	if err := json.NewDecoder(r.Body).Decode(&consumer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateConsumer(consumer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeleteConsumer(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	if err := dao.DeleteConsumer(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}
