package search

import (
	"net/http"
)

//Search writes and sends request to 3rd party API based on given params
func Search(w http.ResponseWriter, r *http.Request) {
	// ps := httprouter.ParamsFromContext(r.Context())
	ConnectDB()

}
