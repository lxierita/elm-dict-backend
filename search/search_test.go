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
	expectedType := "application/json; charset=utf-8"

	if got := res.Header.Get("Content-Type"); got != expectedType {
		t.Errorf("\nunexpected response type from server: \nexpect: %v\ngot: %v", expectedType, got)
	}
}
