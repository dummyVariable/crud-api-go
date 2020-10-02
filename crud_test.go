package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_for_create_handler(t *testing.T) {
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

func Test_for_read_handler(t *testing.T) {
	tests := []struct {
		id      string
		message string
		status  int
	}{
		{"1", "first-message", http.StatusOK},
		{"2", "second-message", http.StatusOK},
		{"5", "No message found", http.StatusNotFound},
	}
	for _, testcase := range tests {
		url := "/" + testcase.id
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		readHandler(w, r)

		resp := w.Result()
		var item Item
		json.Unmarshal(w.Body.Bytes(), &item)

		if resp.StatusCode != testcase.status {
			t.Errorf("error status expecting %v got %v at id %v", testcase.status, resp.StatusCode, testcase.id)
		}
		if item.Message != testcase.message {
			t.Errorf("error status expecting %v got %v at id %v", testcase.status, item.Message, testcase.id)
		}
	}
}
