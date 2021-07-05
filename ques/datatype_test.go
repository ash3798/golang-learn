package main

import (
	"testing"
)

type TestDatatypeInput struct {
	testcase string
	input    interface{}
	expected string
}

func TestDetectDatatype(t *testing.T) {
	dataItems := []TestDatatypeInput{
		{"string input", "teststring", "string"},
		{"int input", 10, "int"},
		{"float input", 20.10, "float64"},
		{"bool", true, "can't figure out type"},
	}

	for _, item := range dataItems {
		result := detectDatatype(item.input)

		if result != item.expected {
			t.Errorf("[%s] returned wrong datatype. Expected : %s ,  got : %s", item.testcase, item.expected, result)
		}
	}
}
