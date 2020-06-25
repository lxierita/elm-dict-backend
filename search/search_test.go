package search

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearch_should_handle_routes_that_contain_the_word_search(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Search))
	defer ts.Close()

	strs := []string{ts.URL, "/search/apple"}
	url := strings.Join(strs, "")

	fmt.Println("New Server URL: \n", url)

	res, err := http.Get(url)

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
