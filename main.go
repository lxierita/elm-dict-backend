package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Search contacts Merriam-Webster API for the definition of inputted word
func Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/search/:word", Search)

	log.Fatal(http.ListenAndServe(":8080", router))
}
