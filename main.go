package main

import (
	"log"
	"net/http"

	routing "lxierita/elm-dict/router"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.HandleOPTIONS = true
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "GET")
		header.Set("Access-Control-Allow-Origin", "127.0.0.1:8000")

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	routing.ConfigRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
