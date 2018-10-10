package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	input, output := "http://localhost:9090/", "helloworld"
	resp, err := http.Get(input)

	if err != nil {
		t.Fatal(
			"For Input:", input,
			"expected:", output,
			"got:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if stringBody := string(body); err != nil || string(body) != output {
		t.Fatal(
			"For input: ", input,
			"expected: ", output,
			"got:", stringBody)
	}

}
