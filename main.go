package main

import "net/http"

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		readHandler(w, r)
	case "POST":
		createHandler(w, r)
	case "PUT":
		updateHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)

	}
}

func main() {
	startDB()
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8000", nil)
}
