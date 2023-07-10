package ch04

import "github.com/bahodge/data_structures_the_fun_way/ch03"

type ArrayStack struct {
	capacity int
	top      int
	values   []int
}

func (as *ArrayStack) Push(value int) {
	if as.capacity == 0 {
		as.values = make([]int, 1)
		as.capacity = cap(as.values)
	}

	if as.top == as.capacity-1 {
		as.values = ch03.ArrayDouble(as.values)
		as.capacity = cap(as.values)
	}
	as.top++
	as.values[as.top] = value
}

func (as *ArrayStack) Pop() int {
	var value int
	if as.top > -1 {
		value = as.values[as.top]
		as.top--
	}

	return value

}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}
