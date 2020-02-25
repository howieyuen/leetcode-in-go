package _707_design_linked_list

type MyLinkedList struct {
	ele []int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{ele: []int{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(this.ele) {
		return -1
	}
	return this.ele[index]
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	reverse(this.ele, 0, len(this.ele)-1)
	this.ele = append(this.ele, val)
	reverse(this.ele, 0, len(this.ele)-1)
}

func reverse(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.ele = append(this.ele, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == len(this.ele) {
		this.AddAtTail(val)
	} else if index > 0 && index < len(this.ele) {
		reverse(this.ele, index, len(this.ele)-1)
		this.AddAtTail(val)
		reverse(this.ele, index, len(this.ele)-1)
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < len(this.ele) {
		this.ele = append(this.ele[0:index], this.ele[index+1:]...)
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
