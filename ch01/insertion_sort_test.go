package ch01

import (
	"reflect"
	"testing"
)

func TestSortedArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := insertionSort(input)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestUnsortedArray(t *testing.T) {
	input := []int{5, 2, 7, 6, 3, 8, 9, 1, 10, 4}
	got := insertionSort(input)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestDuplicates(t *testing.T) {
	input := []int{1, 3, 5, 7, 9, 9, 7, 5, 3, 1}
	got := insertionSort(input)
	want := []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}
