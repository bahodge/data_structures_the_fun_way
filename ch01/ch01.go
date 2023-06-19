package ch01

import "fmt"

func insertionSort(a []int) []int {
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

func stringEqual(s string, t string) bool {
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

func Run() {
	fmt.Println("----------------ch01 start-----------------")
	fmt.Println("InsertionSort 1", insertionSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println("InsertionSort 2", insertionSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println("InsertionSort 3", insertionSort([]int{5, 2, 7, 6, 3, 8, 9, 1, 10, 4}))

	fmt.Println("StringEqual 1", stringEqual("hello", "hello"))
	fmt.Println("StringEqual 2", stringEqual("helLo", "hello"))
	fmt.Println("StringEqual 3", stringEqual("hello ", " hello"))
	fmt.Println("StringEqual 3", stringEqual("hello ", "hello "))
	fmt.Println("----------------ch01 end-----------------")

}
