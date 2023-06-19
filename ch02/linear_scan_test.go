package ch02

import (
	"math"
	"reflect"
	"testing"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func TestLinearScan(t *testing.T) {
	a := fixtures.SortedInts
	target := fixtures.SortedInts[3]
	got := linearScan(a, target)
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestLinearScanNotFound(t *testing.T) {
	a := fixtures.SortedInts
	target := math.MaxInt
	got := linearScan(a, target)
	want := -1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
