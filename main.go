package main

import (
	"Labora_Wallet/utils"
	"Labora_wallet/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//ruta principal
	router.HandleFunc("/", routes.UserMessage)
	//ruta post
	router.HandleFunc("/search", routes.SearchRequest).Methods("POST")
	//ruta get
	router.HandleFunc("/get", utils.SearchResult).Methods("GET")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
