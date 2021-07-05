package main

import "testing"

func TestSwap1(t *testing.T) {
	a := 10
	b := 20

	result := swap1(a, b)

	if result[0] != b || result[1] != a {
		t.Errorf("Values not swapped properly. Expected : a = %d , b = %d and got : a = %d , b = %d", b, a, result[0], result[1])
	} else {
		t.Logf("success. Got :  a = %d , b= %d ", result[0], result[1])
	}
}

func TestSwap2(t *testing.T) {
	a := 10
	b := 20

	result := swap2(a, b)

	if result[0] != b || result[1] != a {
		t.Errorf("Values not swapped properly. Expected : a = %d , b = %d and got : a = %d , b = %d", b, a, result[0], result[1])
	} else {
		t.Logf("success. Got :  a = %d , b= %d ", result[0], result[1])
	}
}
