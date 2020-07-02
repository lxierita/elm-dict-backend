package search

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//MWEndpoint is the endpoint of Merriam-Webster API
const MWEndpoint = "https://www.dictionaryapi.com/api/v3/references/learners/json/"

//APIKey is the api key
const APIKey = "24375962-78c5-4fbc-a585-b37ed4088caf"

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	url := fmt.Sprintf("%v%v?key=%v", MWEndpoint, ps.ByName("word"), APIKey)

	fmt.Fprintf(w, "hello, %s!\n", url)
}
