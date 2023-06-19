package ch01

import (
	"reflect"
	"testing"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func TestSortedArray(t *testing.T) {
	input := fixtures.SortedInts
	got := insertionSort(input)
	want := fixtures.SortedInts

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestUnsortedArray(t *testing.T) {
	input := fixtures.UnsortedInts
	got := insertionSort(input)
	want := fixtures.SortedInts

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestDuplicates(t *testing.T) {
	input := fixtures.UnsortedDupeInts
	got := insertionSort(input)
	want := fixtures.SortedDupeInts

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}
