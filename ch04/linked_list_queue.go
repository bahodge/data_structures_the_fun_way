package ch04

type LinkedListQueue struct {
	front *linkedListQueueNode
	back  *linkedListQueueNode
	size  int
}

type linkedListQueueNode struct {
	value int
	next  *linkedListQueueNode
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (llq *LinkedListQueue) Enqueue(value int) {
	node := &linkedListQueueNode{
		value: value,
	}

	if llq.back == nil {
		llq.front = node
		llq.back = node
	} else {
		llq.back.next = node
		llq.back = node
	}
	llq.size++
}

func (llq *LinkedListQueue) Dequeue() int {
	if llq.front == nil {
		return -1
	}

	val := llq.front.value
	llq.front = llq.front.next
	llq.size--

	if llq.front == nil {
		llq.back = nil
		return -1
	}

	return val

}
