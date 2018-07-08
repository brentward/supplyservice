package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "supplyservice/models"
	. "supplyservice/dao"
	"github.com/globalsign/mgo/bson"
)

var dao = DataObject{}

func GetProduct(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	product, err := dao.FindProductById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, product)
}

func GetProducts(w http.ResponseWriter, _ *http.Request)  {
	products, err := dao.FindAllProducts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.ID = bson.NewObjectId()
	if err := dao.InsertProduct(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateProduct(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeleteProduct(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

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

func init() {
	dao.Server = "localhost"
	dao.Database = "supplyservice"
	dao.Connect()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products", UpdateProduct).Methods("PATCH")
	router.HandleFunc("/products", GetProduct).Methods("GET")
	router.HandleFunc("/products", DeleteProduct).Methods("DELETE")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
