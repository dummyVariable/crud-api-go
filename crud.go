package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func jsonResponseBuilder(i Item) []byte {
	output, _ := json.MarshalIndent(&i, "", "\t")
	return output
}

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
		w.Write(jsonResponseBuilder(payload))
	} else {
		w.WriteHeader(http.StatusOK)
		payload := Item{message}
		w.Write(jsonResponseBuilder(payload))
	}

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	startDB()
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/"))

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var item Item
	json.Unmarshal(body, &item)

	res := updateItem(id, item.Message)

	if res != "Updated" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		i := Item{"error happend"}
		w.Write(jsonResponseBuilder(i))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponseBuilder(Item{res}))
	}
	return
}
