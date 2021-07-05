package main

import "testing"

func TestFindSumOfSquares(t *testing.T) {
	expected := 55

	result := FindSumOfSquares()

	if result != expected {
		t.Errorf("expected :%d , got : %d", expected, result)
	} else {
		t.Logf("success . Got result : %d ", result)
	}
}
