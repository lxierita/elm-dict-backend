package search

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//MWEndpoint is the endpoint of Merriam-Webster API
const MWEndpoint = "/api/v3/references/learners/json/"

//APIKey is the api key
const APIKey = "24375962-78c5-4fbc-a585-b37ed4088caf"

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	r.URL.Scheme = "https"
	r.URL.Host = "dictionaryapi.com"
	r.URL.Path = fmt.Sprintf("%v%v?key=%v", MWEndpoint, ps.ByName("word"), APIKey)

	client := &http.Client{}
	res, err := client.Get(r)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	if body, err := res.Body

}
