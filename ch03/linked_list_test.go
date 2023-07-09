package ch03

import (
	"testing"
)

func TestLinkedListLookup(t *testing.T) {
	n0 := &linkedListNode{
		value: 256,
		next:  nil,
	}
	got := linkedListLookup(n0, 0)
	want := n0

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestLinkedListLookupDoesNotMatch(t *testing.T) {
	n0 := &linkedListNode{
		value: 256,
		next:  nil,
	}

	got := linkedListLookup(n0, 7)

	if got != nil {
		t.Errorf("got %+v, want %+v", got, nil)
	}
}

// This test will add a new node 100_000 times, so the linked list becomes quite large.
// This should theoretically take a long time to traverse
func TestLinkedListLookupLong(t *testing.T) {
	n0 := &linkedListNode{
		value: 256,
		next:  nil,
	}

	curr := n0
	for i := 0; i < 100_000; i++ {
		next := &linkedListNode{
			value: i,
			next:  nil,
		}
		curr.next = next
		curr = next
	}

	got := linkedListLookup(n0, 100_001)

	if got != nil {
		t.Errorf("got %+v, want %+v\n", got, nil)
	}
}

func TestInsertLinkedListAppend(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &linkedListNode{
		value: 1,
		next:  nil,
	}

	linkedListAppend(n0, n1)

	got := linkedListLookup(n0, 1)
	want := n1

	if got != n1 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertHead(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &linkedListNode{
		value: 1,
		next:  nil,
	}

	got, err := linkedListInsert(n0, 0, n1)
	if err != nil {
		t.Errorf("got error %+v\n", err)
	}

	want := n1

	if got != n1 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertMid(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &linkedListNode{
		value: 1,
		next:  nil,
	}
	n2 := &linkedListNode{
		value: 2,
		next:  nil,
	}

	linkedListInsert(n0, 1, n1)

	linkedListInsert(n0, 1, n2)
	got := linkedListLookup(n0, 1)
	want := n2

	if got != n2 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertError(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &linkedListNode{
		value: 1,
		next:  nil,
	}

	_, err := linkedListInsert(n0, 1, n1)

	if err != nil {
		t.Errorf("expected error\n")
	}
}

func TestLinkedListDeleteHead(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}

	got, err := linkedListDelete(n0, 0)

	if err != nil {
		t.Errorf("unexpected error %s\n", err)
	}

	if got != nil {
		t.Errorf("expected nil\n")
	}
}

func TestLinkedListDeleteIndexOutOfBounds(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}

	_, err := linkedListDelete(n0, 10)

	if err == nil {
		t.Errorf("expected error %s\n", err)
	}
}

func TestLinkedListDelete(t *testing.T) {
	n0 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &linkedListNode{
		value: 0,
		next:  nil,
	}
	n2 := &linkedListNode{
		value: 0,
		next:  nil,
	}

	linkedListInsert(n0, 1, n1)
	linkedListInsert(n0, 2, n2)

	head, _ := linkedListDelete(n0, 1)

	if head != n0 {
		t.Errorf("expected head to be n0\n")
	}

	if linkedListLookup(n0, 2) != nil {
		t.Errorf("expected index 2 to be deleted\n")
	}
}
