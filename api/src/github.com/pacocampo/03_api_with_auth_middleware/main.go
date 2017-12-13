package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pacocampo/03_api_with_auth_middleware/auth"
	"github.com/pacocampo/03_api_with_auth_middleware/controllers"
	"github.com/pacocampo/03_api_with_auth_middleware/db"
)

func main() {
	session := db.GetSession()
	uc := controllers.NewUserController(session)
	mux := mux.NewRouter()
	mux.HandleFunc("/auth", auth.Auth)
	mux.HandleFunc("/users", auth.ValidateMiddleware(uc.List)).Methods("GET")
	mux.HandleFunc("/users/{id}", auth.ValidateMiddleware(uc.Get)).Methods("GET")
	mux.HandleFunc("/users/{id}/update", auth.ValidateMiddleware(uc.Update)).Methods("POST")
	mux.HandleFunc("/test", auth.ValidateMiddleware(uc.TestEndpoint)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
