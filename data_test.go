package main

import "testing"

func Test_data_create(t *testing.T) {
	tests := []struct {
		ID   int
		item string
	}{
		{1, "test1"},
		{2, "test2"},
	}

	for _, testcase := range tests {
		createItem(testcase)
	}

	if len(data) != 2 {
		t.Errorf("Expected data items to be 2; Got : %v", len(data))
	}

	if data[0].ID != 1 && data[0].Message != "test1" {
		t.Errorf("expected data to be {1, test1} got:%v", data[0])
	}

	if data[1].ID != 2 && data[1].Message != "test2" {
		t.Errorf("expected data to be {2, test2} got:%v", data[1])
	}
}
