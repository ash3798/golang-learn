package main

func swap1(a int, b int) []int {
	a, b = b, a
	return []int{a, b}
}

func swap2(a int, b int) []int {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return []int{a, b}
}
