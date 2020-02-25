package _155_min_stack

import (
	`math`
)

type MinStack struct {
	data  []int
	index int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:  []int{},
		index: -1,
		min:   math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.index++
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if this.index == -1 {
		return
	} else if this.index == 0 {
		this.min = math.MaxInt64
	} else {
		if this.min == this.data[this.index] {
			min := math.MaxInt64
			for i := 0; i < this.index; i++ {
				if this.data[i] < min {
					min = this.data[i]
				}
			}
			this.min = min
		}
	}
	this.data = this.data[:this.index]
	this.index--
}

func (this *MinStack) Top() int {
	if this.index == -1 {
		return math.MaxInt64
	}
	return this.data[this.index]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
