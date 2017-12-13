package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pacocampo/01_api/controllers"
	"github.com/pacocampo/01_api/db"
)

func main() {
	session := db.GetSession()
	uc := controllers.NewUserController(session)
	mux := mux.NewRouter()
	mux.HandleFunc("/users", uc.List).Methods("GET")
	mux.HandleFunc("/users/create", uc.Create).Methods("POST")
	mux.HandleFunc("/users/{id}", uc.Get).Methods("GET")
	mux.HandleFunc("/users/{id}/update", uc.Update).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
