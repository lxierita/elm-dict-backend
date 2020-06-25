package main_test

import (
	"fmt"
	"io/ioutil"
	"log"
	. "main"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch_handler_should_handle_word_search(t *testing.T) {
	ts := httptest.NewServer(Search())
	defer ts.Close()

	ts.URL += "search/apple"

	res, err := http.Get(ts.URL)
	fmt.Println(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", result)
}
