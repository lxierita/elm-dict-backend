package search

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("word"))
}
