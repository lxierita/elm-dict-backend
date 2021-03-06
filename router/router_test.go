package router_test

import (
	"fmt"
	"io/ioutil"
	routing "lxierita/elm-dict/router"
	"net/http"
	"net/http/httptest"
	"testing"

	hr "github.com/julienschmidt/httprouter"
)

func TestRouter_should_register_handlers_correctly(t *testing.T) {

	//implement routing
	router := hr.New()
	routing.ConfigRoutes(router)
	searchRoute := "/test/:name"

	//add stubbed handler
	router.Handle(http.MethodGet, searchRoute, func(w http.ResponseWriter, req *http.Request, ps hr.Params) {
		fmt.Fprintf(w, "%v!", ps.ByName("name"))
	})

	//serve requests
	req := httptest.NewRequest(http.MethodGet, "/test/boo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Errorf("ioutil.ReadAll(body) = %v", err)
	}

	if want := "boo!"; string(body) != want {
		t.Errorf("expected %v ; got %v", want, string(body))
	}

}
