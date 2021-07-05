package main

import "testing"

func TestFindPermutations(t *testing.T) {
	a := "abc"

	FindPermutations([]rune(a))

	t.Errorf("failed the test to check output")
}
