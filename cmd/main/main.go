package main

import (
	"log"
	"net/http"

	"github.com/HaydnMeyburgh/booking-management-system/pkg/config"
	"github.com/HaydnMeyburgh/booking-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.DBConnection()

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
