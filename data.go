package main

//Item model for api
type Item struct {
	ID      int
	Message string
}

var data []Item

func createItem(i Item) {
	data = append(data, i)
}
