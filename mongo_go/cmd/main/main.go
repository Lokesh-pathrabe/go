package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lokesh/mongo_go/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
