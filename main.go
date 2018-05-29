package main

import (
	"net/http"
	"log"
	"widgetFactory/service"

	"github.com/gorilla/mux"
)

const (
	port = ":8000"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/authenticate", service.CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/widgets", service.ValidateMiddleware(service.GetWidgets)).Methods("GET")
	router.HandleFunc("/users", service.ValidateMiddleware(service.GetUsers)).Methods("GET")
	router.HandleFunc("/widgets/{id}", service.ValidateMiddleware(service.GetWidget)).Methods("GET")
	router.HandleFunc("/users/{id}", service.ValidateMiddleware(service.GetUser)).Methods("GET")
	router.HandleFunc("/widgets", service.ValidateMiddleware(service.CreateWidget)).Methods("POST")
	router.HandleFunc("/widgets/{id}", service.ValidateMiddleware(service.UpdateWidget)).Methods("PUT")
	log.Fatal(http.ListenAndServe(port, router))
}