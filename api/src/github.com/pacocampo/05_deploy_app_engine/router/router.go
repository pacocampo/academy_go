package router

import (
	"github.com/gorilla/mux"
	"github.com/pacocampo/05_deploy_app_engine/auth"
	"github.com/pacocampo/05_deploy_app_engine/controllers"
	"github.com/pacocampo/05_deploy_app_engine/db"
)

func New() *mux.Router {
	session := db.GetSession()
	uc := controllers.NewUserController(session)
	mux := mux.NewRouter()
	mux.HandleFunc("/auth", auth.Auth)
	mux.HandleFunc("/users", auth.ValidateMiddleware(uc.List)).Methods("GET")
	mux.HandleFunc("/users/create", auth.ValidateMiddleware(uc.Create)).Methods("POST")
	mux.HandleFunc("/users/{id}", auth.ValidateMiddleware(uc.Get)).Methods("GET")
	mux.HandleFunc("/users/{id}/update", auth.ValidateMiddleware(uc.Update)).Methods("POST")
	mux.HandleFunc("/test", auth.ValidateMiddleware(uc.TestEndpoint)).Methods("GET")

	return mux
}
