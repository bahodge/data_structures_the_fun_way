package ch04

type LinkedListStack struct {
	head *linkedListStackNode
	size int
}

type linkedListStackNode struct {
	value int
	next  *linkedListStackNode
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (lls *LinkedListStack) Push(value int) {

	newNode := &linkedListStackNode{
		value: value,
		next:  lls.head,
	}
	lls.head = newNode
	lls.size++
}

func (lls *LinkedListStack) Pop() int {
	if lls.head != nil {
		value := lls.head.value
		lls.head = lls.head.next
		lls.size--
		return value
	}

	return -1
}
