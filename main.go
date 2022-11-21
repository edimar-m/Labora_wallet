package main

import (
	"log"
	"net/http"

	"github.com/edimar-m/Labora_wallet/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//ruta principal
	router.HandleFunc("/", controllers.UserMessage)
	//ruta post
	router.HandleFunc("/search", controllers.SearchRequest).Methods("POST")
	//ruta get
	router.HandleFunc("/get", controllers.SearchResult).Methods("GET")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
