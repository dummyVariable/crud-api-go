package main

//Item model for api
type Item struct {
	Message string `json:"message"`
}

var data map[int]string
var counter int

func startDB() {
	data = make(map[int]string)
	createItem("first-message")
	createItem("second-message")

}

func createItem(message string) {
	counter++
	data[counter] = message
}

func readItem(id int) string {
	item := data[id]
	if item == "" {
		return "None"
	}
	return item
}

func updateItem(id int, message string) string {
	_, present := data[id]
	if !present {
		return "None"
	}
	if message == "" {
		return "None"
	}
	data[id] = message
	return "Updated"
}

func deleteItem(id int) string {
	_, present := data[id]
	if !present {
		return "None"
	}
	delete(data, id)
	return "Deleted"
}
