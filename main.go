package main

import (
	"log"
	"net/http"

	routing "lxierita/elm-dict/router"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	routing.ConfigRoutes(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
