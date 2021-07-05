package main

import "fmt"

func FindPermutations(a []rune) {
	permute(a, 0)
}

func permute(a []rune, i int) {
	if i > len(a) {
		fmt.Println(string(a))
		return
	}

	permute(a, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
