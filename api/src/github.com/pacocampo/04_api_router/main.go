package main

import (
	"log"
	"net/http"

	"github.com/pacocampo/04_api_router/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.New()))
}
