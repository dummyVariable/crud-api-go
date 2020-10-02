package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_for_create(t *testing.T) {
	tests := []struct {
		message io.Reader
		status  int
	}{
		{strings.NewReader(`{"message" : "first-message"}`), http.StatusOK},
		{strings.NewReader(`{"messag" : "mis-message"}`), http.StatusUnprocessableEntity},
		{strings.NewReader(`{"message" : "second-message"}`), http.StatusOK},
	}
	for _, testcase := range tests {
		r := httptest.NewRequest("POST", "/", testcase.message)
		w := httptest.NewRecorder()
		createHandler(w, r)

		resp := w.Result()
		var body []byte
		resp.Body.Read(body)
		if resp.StatusCode != testcase.status {
			t.Errorf("Error at creating item expected :%v got %v %v", testcase.status, resp.StatusCode, body)
		}

	}

}
