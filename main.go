package main

import (
	"fmt"
	"log"
	. "lxierita/elm-dict/search"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Index prints out a welcome message
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.Handler("GET", "/search/:word", http.HandlerFunc(Search))

	log.Fatal(http.ListenAndServe(":8080", router))
}
