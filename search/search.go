package search

import (
<<<<<<< HEAD
	"fmt"
	"io/ioutil"
	"log"
=======
>>>>>>> middleware
	"net/http"
)

//MWEndpoint is the endpoint of Merriam-Webster API
const MWEndpoint = "/api/v3/references/learners/json/"

//APIKey is the api key
const APIKey = "24375962-78c5-4fbc-a585-b37ed4088caf"

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	// ps := httprouter.ParamsFromContext(r.Context())
	
	r.URL.Scheme = "https"
	r.URL.Host = "dictionaryapi.com"
	r.URL.Path = fmt.Sprintf("%v%v?key=%v", MWEndpoint, ps.ByName("word"), APIKey)

	ts := &http.Transport{
		Proxy: http.ProxyURL(r.URL),
	}
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	res, err := ts.RoundTrip(r)
	if err != nil {
		log.Fatalf("RoundTrip failed: \n%v", err)
	}
	body, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		log.Fatalf("Couldn't read response: \n %v", err1)
	}

	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(body)
	if err2 != nil {
		log.Fatalf("ResponseWriter failed to write: \n %v", err2)
	}
	ts.CloseIdleConnections()
}
