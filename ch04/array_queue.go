package ch04

import (
	"github.com/bahodge/data_structures_the_fun_way/ch03"
)

type ArrayQueue struct {
	front    int
	back     int
	queue    []int
	capacity int
	length   int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}

}

func (aq *ArrayQueue) Enqueue(value int) {

	if aq.capacity == 0 {
		aq.queue = make([]int, 1)
		aq.capacity = cap(aq.queue)
	}

	// if the queue length is at capacity
	if aq.capacity-1 == aq.back {
		aq.queue = ch03.ArrayDouble(aq.queue)
		aq.capacity = cap(aq.queue)
	}

	aq.queue[aq.back] = value
	aq.back++
	aq.length++

}

func (aq *ArrayQueue) Dequeue() int {
	// we have reached the end of the queue
	// gte because in the Enqueue func we are incrementing the back, so back will always
	// be 1 place past the actual back index
	if aq.front >= aq.back {
		return -1
	}

	// increment the front index
	val := aq.queue[aq.front]
	aq.front++
	aq.length--
	return val
}
