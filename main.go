package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Address struct {
	Street string `json:"street"`
}

type Customer struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

// Customers TODO: replace with DB
var Customers []Customer

func allCustomers(w http.ResponseWriter, _ *http.Request) {
	log.Println("Endpoint Hit: allCustomers")
	err := json.NewEncoder(w).Encode(Customers)
	if err != nil {
		log.Println("It broke")
		return
	}
}

func customer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Customers
	// if the customer.Id equals the key we pass in
	// return the customer encoded as JSON
	for i := range Customers {
		if Customers[i].Id == key {
			err := json.NewEncoder(w).Encode(Customers[i])
			if err != nil {
				log.Println("It broke")
				return
			}
			// no need to continue searching
			break
		}
	}
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/all", allCustomers)
	myRouter.HandleFunc("/customer/{id}", customer)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	log.Println("Rest API v2.0 - Mux Routers")
	Customers = []Customer{
		{Id: "1", Name: "Person People", Address: Address{Street: "somestreet"}},
	}
	handleRequests()
}
