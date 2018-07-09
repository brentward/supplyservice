package controller

import (
	. "supplyservice/models"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
)

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
	params := mux.Vars(r)
	//// Read body
	//b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	//	return
	//}
	defer r.Body.Close()
	var product Product
	//var product_json interface{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
	//if err := json.Unmarshal(b, &product_json); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	//product := product_json.(map[string]interface{})
	if err := dao.UpdateProduct(params["id"], product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	if err := dao.DeleteProduct(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}
