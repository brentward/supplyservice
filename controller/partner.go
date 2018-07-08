package controller

import (
	. "supplyservice/models"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
)

func GetPartner(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	partner, err := dao.FindPartnerById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, partner)
}

func GetPartners(w http.ResponseWriter, _ *http.Request)  {
	partners, err := dao.FindAllPartners()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, partners)
}

func CreatePartner(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var partner Partner
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	partner.ID = bson.NewObjectId()
	if err := dao.InsertPartner(partner); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, partner)
}

func UpdatePartner(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var partner Partner
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdatePartner(partner); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeletePartner(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var partner Partner
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeletePartner(partner); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}
