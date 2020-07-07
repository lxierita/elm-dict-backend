package router

import (
	"fmt"
	s "lxierita/elm-dict/search"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Index prints out a welcome message
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//ConfigRoutes configures the routes and handlers of the app
func ConfigRoutes(r *httprouter.Router) {
	r.GET("/", Index)
	r.Handler("GET", "/search/:word", allowAll(http.HandlerFunc(s.Search)))
}

func allowAll(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
