package search

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch_should_send_request_to_mw_api(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/search/boo", nil)
	Search(w, req)

	res := w.Result()

	if got := res.StatusCode; got != 200 {
		t.Errorf("\nBad status code: \nexpect: 200\ngot: %v", got)
	}
}
