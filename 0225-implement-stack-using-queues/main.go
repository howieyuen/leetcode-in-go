package _225_implement_stack_using_queues

type Node struct {
	value int
	next  *Node
}

type MyStack struct {
	head *Node
}

func Constructor() MyStack {
	return MyStack{nil}
}

func (this *MyStack) Push(x int) {
	old := this.head
	this.head = &Node{x, old}
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	top := this.head.value
	this.head = this.head.next
	return top
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.head.value
}

func (this *MyStack) Empty() bool {
	return this.head == nil
}
