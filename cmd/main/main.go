package main

import (
	"cyrusbesabella/go-bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010\n")
	log.Fatal(http.ListenAndServe(":9010", r))
}

