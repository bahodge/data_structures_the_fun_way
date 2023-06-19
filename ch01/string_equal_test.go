package ch01

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func TestEqual(t *testing.T) {
	str := fixtures.LongString
	test := fixtures.LongString
	got := stringEqual(str, test)
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestNotEqualLength(t *testing.T) {
	str := fixtures.LongString
	test := fixtures.LongString + " "
	got := stringEqual(str, test)
	want := false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestNotEqual(t *testing.T) {
	str := fixtures.LongString
	test := strings.ToUpper(fixtures.LongString)
	got := stringEqual(str, test)
	want := false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}
