package search

import (
	"net/http"
)

//MWEndpoint is the endpoint of Merriam-Webster API
const MWEndpoint = "/api/v3/references/learners/json/"

//APIKey is the api key
const APIKey = "24375962-78c5-4fbc-a585-b37ed4088caf"

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	// ps := httprouter.ParamsFromContext(r.Context())

	// r.URL.Host = "dictionaryapi.com"
	// r.URL.Path = fmt.Sprintf("%v%v?key=%v", MWEndpoint, ps.ByName("word"), APIKey)

	// ts := &http.Transport{}
	// ts.Proxy(r)

	w.WriteHeader(http.StatusOK)
}
