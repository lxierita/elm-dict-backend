package search

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSearch_should_handle_routes_that_contain_the_search_param(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search/apple", nil)
	Search(w, req)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Errorf("ioutil.ReadAll(body) = %v", err)
	}
	want := "hello apple!"

	fmt.Printf("response: %s\n", body)

	if string(body) != want {
		t.Errorf("expected %v, got %v", want, string(body))
	}
}
