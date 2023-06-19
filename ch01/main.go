package main

import "fmt"

func InsertionSort(a []int) []int {
	for i := range a {
		current := a[i]
		j := i - 1
		for j >= 0 && a[j] > current {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = current
		i = i + 1
	}

	return a
}

func StringEqual(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tRunes := []rune(t)
	for i, n := range s {
		if n != tRunes[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("1", InsertionSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println("2", InsertionSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println("3", InsertionSort([]int{5, 2, 7, 6, 3, 8, 9, 1, 10, 4}))
}
