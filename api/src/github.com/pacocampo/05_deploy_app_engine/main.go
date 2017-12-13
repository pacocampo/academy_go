package main

import (
	"log"
	"net/http"

	"github.com/pacocampo/05_deploy_app_engine/router"
)

func init() {
	// r := router.New()
	// r.Handle("/", http.FileServer(http.Dir(".")))
	// r.Handle("/", r)
}

func main() {
	r := router.New()
	r.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", router.New()))
}
