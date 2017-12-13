package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	eVar := os.Getenv("SECRET_WORD")
	fmt.Fprintln(w, "Hello gopher from docker <3 y con Mux (: y "+eVar)
}
