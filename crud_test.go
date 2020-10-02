package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_for_create(t *testing.T) {
	tests := []struct {
		message io.Reader
		status  string
	}{
		{strings.NewReader("data=first-message"), "created"},
		{strings.NewReader("dat=mis-message"), "error at creating"},
		{strings.NewReader("data=second-message"), "created"},
	}
	for _, testcase := range tests {
		r := httptest.NewRequest("POST", "/", testcase.message)
		w := httptest.NewRecorder()
		crudHandler(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error at creating item")
		}
		if string(body) != testcase.status {
			t.Errorf("Error at creating item")
		}
	}

}
