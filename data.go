package main

//Item model for api

var data map[int]string
var counter int

func startDB() {
	data = make(map[int]string)
}

func createItem(message string) {
	counter++
	data[counter] = message
}
