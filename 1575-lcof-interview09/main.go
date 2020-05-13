package _575_lcof_interview09

type CQueue struct {
	input  stack
	output stack
}

type stack []int

func (s *stack) pop() int {
	n := len(*s)
	ans := (*s)[n-1]
	*s = (*s)[:n-1]
	return ans
}

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func Constructor() CQueue {
	return CQueue{
		input:  nil,
		output: nil,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.input.push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.output) != 0 {
		return this.output.pop()
	}
	if len(this.input) == 0 {
		return -1
	}
	for len(this.input) != 0 {
		this.output.push(this.input.pop())
	}
	return this.output.pop()
}
