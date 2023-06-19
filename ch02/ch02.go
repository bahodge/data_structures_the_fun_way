package ch02

import (
	"fmt"
	"math"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func linearScan(a []int, target int) int {
	for i, n := range a {
		if n == target {
			return i
		}
	}
	return -1
}

func binarySearch(a []int, target int) int {
	i := 0
	j := len(a) - 1
	for i <= j {
		mid := int(math.Floor(float64(i+j)) / 2)

		if a[mid] == target {
			return mid
		}

		if a[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return -1
}

func Run() {
	fmt.Println("----------------ch02 start-----------------")
	fmt.Println("linearScan 1", linearScan(fixtures.SortedInts, 3))
	fmt.Println("linearScan 2", linearScan(fixtures.SortedDupeInts, 8))
	fmt.Println("linearScan 2", linearScan(fixtures.SortedDupeInts, 999))

	fmt.Println("binarySearch 1", binarySearch(fixtures.SortedInts, 3))
	fmt.Println("binarySearch 2", binarySearch(fixtures.SortedDupeInts, 8))
	fmt.Println("binarySearch 2", binarySearch(fixtures.SortedDupeInts, 999))
	fmt.Println("----------------ch02 end-----------------")

}
