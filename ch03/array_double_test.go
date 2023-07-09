package ch03

import (
	"testing"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func TestArrayDouble(t *testing.T) {
	a := fixtures.SortedInts
	got := arrayDouble(a)
	// Want this capacity
	want := len(a) * 2

	if cap(got) != want {
		t.Errorf("got %+v, want %+v", cap(got), want)
	}
}
