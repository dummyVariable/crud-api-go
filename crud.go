package main

import (
	"encoding/json"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	startDB()
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var item Item
	json.Unmarshal(body, &item)

	if item.Message == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		createItem(item.Message)
		w.WriteHeader(http.StatusOK)
	}
	return
}
