package main

import "testing"

func Test_data_create(t *testing.T) {
	startDB()

	tests := []string{
		"test1",
		"test2",
	}

	for i, testcase := range tests {
		createItem(testcase)
		id := i + 1
		if data[id] != testcase {
			t.Errorf("expected data to be test1 got:%v", data[id])
		}
	}

	if len(data) != 2 {
		t.Errorf("Expected data items to be 2; Got : %v", len(data))
	}
}

func Test_for_read(t *testing.T) {
	tests := []string{
		"test1",
		"test2",
		"None",
	}

	for i, testcase := range tests {
		id := i + 1
		got := readItem(id)
		if got != testcase {
			t.Errorf("Expected %v but got %v", testcase, got)
		}
	}
}
