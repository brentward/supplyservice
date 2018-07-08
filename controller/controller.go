package controller

import (
	"net/http"
	"encoding/json"
	. "supplyservice/dao"
)

var dao = DataObject{}

func init() {
	dao.Server = "localhost"
	dao.Database = "supplyservice"
	dao.Connect()
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
