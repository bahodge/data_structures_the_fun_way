package ch04

import (
	"fmt"
)

func Run() {
	as := NewArrayStack()
	lls := NewLinkedListStack()
	aq := NewArrayQueue()
	llq := NewLinkedListQueue()

	fmt.Println("----------------ch04 start-----------------")
	as.Push(10)
	as.Push(2)
	as.Push(7)
	fmt.Printf("array stack top: %d, capacity: %d\n", as.top, as.capacity)
	as.Pop()
	as.Pop()
	fmt.Printf("array stack top: %d, capacity: %d\n", as.top, as.capacity)

	lls.Push(10)
	lls.Push(2)
	lls.Push(7)
	fmt.Printf("linked list stack head: %d, size: %d\n", lls.head.value, lls.size)
	lls.Pop()
	lls.Pop()
	fmt.Printf("linked list stack head: %d, size: %d\n", lls.head.value, lls.size)

	aq.Enqueue(10)
	aq.Enqueue(2)
	aq.Enqueue(7)
	fmt.Printf("array queue: front val: %d, front index: %d, length: %d, capacity: %d\n", aq.queue[aq.front], aq.front, aq.length, aq.capacity)
	aq.Dequeue()
	aq.Dequeue()
	fmt.Printf("array queue: front val: %d, front index: %d, length: %d, capacity: %d\n", aq.queue[aq.front], aq.front, aq.length, aq.capacity)

	llq.Enqueue(10)
	llq.Enqueue(2)
	llq.Enqueue(7)
	fmt.Printf("linked list queue: front: %+v, size: %d\n", llq.front, llq.size)
	llq.Dequeue()
	llq.Dequeue()
	fmt.Printf("linked list queue: front: %+v, size: %d\n", llq.front, llq.size)
	fmt.Println("----------------ch04 end-----------------")
}
