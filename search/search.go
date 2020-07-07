package search

import (
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

	test := []byte(ps.ByName("word"))

	w.WriteHeader(http.StatusOK)
	w.Write(test)
}
