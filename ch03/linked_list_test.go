package ch03

import (
	"testing"
)

func TestLinkedListLookup(t *testing.T) {
	n0 := &LinkedListNode{
		value: 256,
		next:  nil,
	}
	got := LinkedListLookup(n0, 0)
	want := n0

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestLinkedListLookupDoesNotMatch(t *testing.T) {
	n0 := &LinkedListNode{
		value: 256,
		next:  nil,
	}

	got := LinkedListLookup(n0, 7)

	if got != nil {
		t.Errorf("got %+v, want %+v", got, nil)
	}
}

// This test will add a new node 100_000 times, so the linked list becomes quite large.
// This should theoretically take a long time to traverse
func TestLinkedListLookupLong(t *testing.T) {
	n0 := &LinkedListNode{
		value: 256,
		next:  nil,
	}

	curr := n0
	for i := 0; i < 100_000; i++ {
		next := &LinkedListNode{
			value: i,
			next:  nil,
		}
		curr.next = next
		curr = next
	}

	got := LinkedListLookup(n0, 100_001)

	if got != nil {
		t.Errorf("got %+v, want %+v\n", got, nil)
	}
}

func TestInsertLinkedListAppend(t *testing.T) {
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &LinkedListNode{
		value: 1,
		next:  nil,
	}

	LinkedListAppend(n0, n1)

	got := LinkedListLookup(n0, 1)
	want := n1

	if got != n1 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertHead(t *testing.T) {
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &LinkedListNode{
		value: 1,
		next:  nil,
	}

	got, err := LinkedListInsert(n0, 0, n1)
	if err != nil {
		t.Errorf("got error %+v\n", err)
	}

	want := n1

	if got != n1 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertMid(t *testing.T) {
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &LinkedListNode{
		value: 1,
		next:  nil,
	}
	n2 := &LinkedListNode{
		value: 2,
		next:  nil,
	}

	LinkedListInsert(n0, 1, n1)

	LinkedListInsert(n0, 1, n2)
	got := LinkedListLookup(n0, 1)
	want := n2

	if got != n2 {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}

func TestLinkedListInsertError(t *testing.T) {
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &LinkedListNode{
		value: 1,
		next:  nil,
	}

	_, err := LinkedListInsert(n0, 1, n1)

	if err != nil {
		t.Errorf("expected error\n")
	}
}

func TestLinkedListDeleteHead(t *testing.T) {
	n0 := &LinkedListNode{
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
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}

	_, err := linkedListDelete(n0, 10)

	if err == nil {
		t.Errorf("expected error %s\n", err)
	}
}

func TestLinkedListDelete(t *testing.T) {
	n0 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n1 := &LinkedListNode{
		value: 0,
		next:  nil,
	}
	n2 := &LinkedListNode{
		value: 0,
		next:  nil,
	}

	LinkedListInsert(n0, 1, n1)
	LinkedListInsert(n0, 2, n2)

	head, _ := linkedListDelete(n0, 1)

	if head != n0 {
		t.Errorf("expected head to be n0\n")
	}

	if LinkedListLookup(n0, 2) != nil {
		t.Errorf("expected index 2 to be deleted\n")
	}
}
