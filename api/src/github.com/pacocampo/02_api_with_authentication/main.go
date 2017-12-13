package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pacocampo/02_api_with_authentication/auth"
	"github.com/pacocampo/02_api_with_authentication/controllers"
	"github.com/pacocampo/02_api_with_authentication/db"
)

func main() {
	session := db.GetSession()
	uc := controllers.NewUserController(session)
	mux := mux.NewRouter()
	mux.HandleFunc("/auth", auth.Auth)
	mux.HandleFunc("/users", uc.List).Methods("GET")
	mux.HandleFunc("/users/{id}", uc.Get).Methods("GET")
	mux.HandleFunc("/users/{id}/update", uc.Update).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
