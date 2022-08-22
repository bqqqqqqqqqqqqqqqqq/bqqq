package main

import (
	"github.com/akhil/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"

	"log"

	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
