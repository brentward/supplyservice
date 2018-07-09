package main

import (
	. "supplyservice/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products", UpdateProduct).Methods("PATCH")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/partners", GetPartners).Methods("GET")
	router.HandleFunc("/partners/{id}", GetPartner).Methods("GET")
	router.HandleFunc("/partners", CreatePartner).Methods("POST")
	router.HandleFunc("/partners", UpdatePartner).Methods("PATCH")
	router.HandleFunc("/partners/{id}", DeletePartner).Methods("DELETE")
	router.HandleFunc("/tickets", GetTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", GetTicket).Methods("GET")
	router.HandleFunc("/tickets", CreateTicket).Methods("POST")
	router.HandleFunc("/tickets", UpdateTicket).Methods("PATCH")
	router.HandleFunc("/tickets/{id}", DeleteTicket).Methods("DELETE")
	router.HandleFunc("/transactions", GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", GetTransaction).Methods("GET")
	router.HandleFunc("/transactions", CreateTransaction).Methods("POST")
	router.HandleFunc("/transactions", UpdateTransaction).Methods("PATCH")
	router.HandleFunc("/transactions/{id}", DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/consumers", GetConsumers).Methods("GET")
	router.HandleFunc("/consumers/{id}", GetConsumer).Methods("GET")
	router.HandleFunc("/consumers", CreateConsumer).Methods("POST")
	router.HandleFunc("/consumers", UpdateConsumer).Methods("PATCH")
	router.HandleFunc("/consumers/{id}", DeleteConsumer).Methods("DELETE")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
