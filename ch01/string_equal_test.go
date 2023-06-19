package ch01

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	str := "hello world"
	test := "hello world"
	got := stringEqual(str, test)
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestNotEqualLength(t *testing.T) {
	str := "hello world"
	test := "hello world "
	got := stringEqual(str, test)
	want := false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func TestNotEqual(t *testing.T) {
	str := "hello world"
	test := "helLo world"
	got := stringEqual(str, test)
	want := false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}
