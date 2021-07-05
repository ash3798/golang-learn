package main

import "testing"

func TestReverseSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	expected := []int{4, 3, 2, 1}

	reverse(nums)

	for i, _ := range nums {
		if nums[i] != expected[i] {
			t.Errorf("reverse failed , expected : %v  ,  got : %v ", expected, nums)
		}
	}

	//if passed through loop that means all values matched
	t.Logf("success.  result : %v ", nums)
}
