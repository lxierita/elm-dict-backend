package search

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Host is the host address
const Host = "dictionaryapi.com"

//MWEndpoint is the endpoint of Merriam-Webster API
const MWEndpoint = "/api/v3/references/learners/json/"

//APIKey is the api key
const APIKey = "24375962-78c5-4fbc-a585-b37ed4088caf"

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	url := fmt.Sprintf("https://%v%v%v?key=%v", Host, MWEndpoint, ps.ByName("word"), APIKey)

	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	body, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		log.Fatalf("Couldn't read response: \n %v", err1)
	}
	_, err2 := w.Write(body)
	if err2 != nil {
		log.Fatalf("ResponseWriter failed to write: \n %v", err2)
	}

}
