package main

import "testing"

func Test_for_create(t *testing.T) {
	startDB()

	tests := []struct {
		ID      int
		message string
	}{
		{1, "test1"},
		{2, "test2"},
	}

	for _, testcase := range tests {
		createItem(testcase.message)
		if data[testcase.ID] != testcase.message {
			t.Errorf("expected data to be test1 got:%v", data[testcase.ID])
		}
	}

	if len(data) != 2 {
		t.Errorf("Expected data items to be 2; Got : %v", len(data))
	}
}

func Test_for_read(t *testing.T) {
	tests := []struct {
		ID      int
		message string
	}{
		{1, "test1"},
		{2, "test2"},
		{3, "None"},
	}

	for _, testcase := range tests {
		got := readItem(testcase.ID)
		if got != testcase.message {
			t.Errorf("Expected %v but got %v", testcase.message, got)
		}
	}
}

func Test_for_update(t *testing.T) {
	tests := []struct {
		ID      int
		message string
		status  string
	}{
		{1, "test1updated", "Updated"},
		{22, "None", "None"},
	}

	for _, testcase := range tests {
		got := updateItem(testcase.ID, testcase.message)

		if got != testcase.status {
			t.Errorf("cannot update the item at ID %v", testcase.ID)
		}

	}
	got := readItem(1)
	if got != "test1updated" {
		t.Errorf("Failed updating at %v expected %v got %v", 1, "test1updated", got)
	}
}

func Test_for_delete(t *testing.T) {
	tests := []struct {
		ID     int
		status string
	}{
		{1, "Deleted"},
		{5, "None"},
	}

	for _, testcase := range tests {
		got := deleteItem(testcase.ID)

		if got != testcase.status {
			t.Errorf("cannot delete item at %v", testcase.ID)
		}
	}
}
