package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func readHandler(w http.ResponseWriter, r *http.Request) {
	startDB()
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))
	message := readItem(id)
	if message == "None" {
		w.WriteHeader(http.StatusNotFound)
		payload := Item{"No message found"}
		output, _ := json.MarshalIndent(&payload, "", "\t")
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusOK)
		payload := Item{message}
		output, _ := json.MarshalIndent(&payload, "", "\t")
		w.Write(output)
	}

}
