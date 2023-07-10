package ch03

import (
	"fmt"

	"github.com/bahodge/data_structures_the_fun_way/fixtures"
)

func ArrayDouble(a []int) []int {
	length := len(a)
	newArr := make([]int, 2*length)
	i := 0
	for i < length {
		newArr[i] = a[i]
		i++
	}

	return newArr
}

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func LinkedListLookup(head *LinkedListNode, index int) *LinkedListNode {
	current := head
	i := 0
	for i < index && current != nil {
		current = current.next
		i++
	}

	return current
}

func LinkedListAppend(prev *LinkedListNode, newNode *LinkedListNode) {
	newNode.next = prev.next
	prev.next = newNode
}

func LinkedListInsert(head *LinkedListNode, index int, newNode *LinkedListNode) (*LinkedListNode, error) {
	if index == 0 {
		newNode.next = head
		return newNode, nil
	}

	current := head
	var previous *LinkedListNode = nil
	count := 0

	for count < index && current != nil {
		previous = current
		current = current.next
		count++
	}

	// we ran off the end of the linked list
	if count < index {
		return nil, fmt.Errorf("linked list length is less than %d\n", index)
	}

	// We have found the index, perform op
	newNode.next = previous.next
	previous.next = newNode

	return head, nil

}

func linkedListDelete(head *LinkedListNode, index int) (*LinkedListNode, error) {
	if head == nil {
		return nil, nil
	}

	if index == 0 {
		newHead := head.next
		head.next = nil
		return newHead, nil
	}

	current := head
	var previous *LinkedListNode = nil
	count := 0

	for count < index && current != nil {
		previous = current
		current = current.next
		count++
	}

	if current != nil {
		previous.next = current.next
		current.next = nil
	} else {
		return nil, fmt.Errorf("invalid index %d\n", index)
	}

	return head, nil

}

func Run() {

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

	LinkedListAppend(n0, n1)
	LinkedListAppend(n1, n2)

	fmt.Println("----------------ch03 start-----------------")
	fmt.Printf("arrayDouble 1 before %d, after %d\n", len(fixtures.SortedInts), len(ArrayDouble(fixtures.SortedInts)))
	fmt.Printf("arrayDouble 2 before %d, after %d\n", len(fixtures.SortedDupeInts), len(ArrayDouble(fixtures.SortedDupeInts)))
	fmt.Printf("linkedListLookup 1 %+v\n", LinkedListLookup(n0, 0))
	fmt.Printf("linkedListLookup 2 %+v\n", LinkedListLookup(n0, 2))
	fmt.Printf("linkedListLookup 3 %+v\n", LinkedListLookup(n0, 8))
	fmt.Println("----------------ch03 end-----------------")

}
