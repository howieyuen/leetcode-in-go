package _705_design_hashset

import (
	`strconv`
)

const (
	defaultCapacity  int     = 16
	growLoadFactor   float32 = 0.75
	shrinkLoadFactor float32 = 0.25
)

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func newNode(val int) *Node {
	return &Node{
		Val:  val,
		Prev: nil,
		Next: nil,
	}
}

type MyHashSet struct {
	bucketArr []*Node
	length    int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		bucketArr: make([]*Node, defaultCapacity),
		length:    0,
	}
}

func (this *MyHashSet) Add(key int) {
	node, ok := this.searchNode(key)
	if ok {
		return
	}
	
	newKey := newNode(key)
	if node != nil {
		node.Next = newKey
		newKey.Prev = node
	} else {
		slot := getSlot(key, cap(this.bucketArr))
		this.bucketArr[slot] = newKey
	}
	
	this.length++
	if this.loadFactor() >= growLoadFactor {
		this.grow()
	}
}

func (this *MyHashSet) Remove(key int) {
	node, ok := this.searchNode(key)
	if !ok {
		return
	}
	
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		slot := getSlot(key, cap(this.bucketArr))
		this.bucketArr[slot] = node.Next
	}
	
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	
	this.length--
	if this.loadFactor() <= shrinkLoadFactor {
		this.shrink()
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, isExisted := this.searchNode(key)
	return isExisted
}

func (this *MyHashSet) searchNode(key int) (*Node, bool) {
	slot := getSlot(key, cap(this.bucketArr))
	node := this.bucketArr[slot]
	for {
		switch {
		case node == nil:
			fallthrough
		case node.Val != key && node.Next == nil:
			return node, false
		case node.Val == key:
			return node, true
		default:
			node = node.Next
		}
	}
}

func (this *MyHashSet) grow() {
	newSize := len(this.bucketArr) * 2
	this.resize(newSize)
}

func (this *MyHashSet) shrink() {
	newSize := len(this.bucketArr) / 2
	if defaultCapacity > newSize {
		newSize = defaultCapacity
	}
	this.resize(newSize)
}

func (this *MyHashSet) loadFactor() float32 {
	return float32(this.length) / float32(cap(this.bucketArr))
}

func (this *MyHashSet) resize(size int) {
	newBucket := make([]*Node, size)
	for i := range this.bucketArr {
		node := this.bucketArr[i]
		var keysInBucket []int
		for {
			if node == nil {
				break
			}
			keysInBucket = append(keysInBucket, node.Val)
			node = node.Next
		}
		
		for j := range keysInBucket {
			key := keysInBucket[j]
			newSlot := getSlot(key, size)
			node = newBucket[newSlot]
			head := newNode(key)
			head.Next = node
			if node != nil {
				node.Prev = head
			}
			newBucket[newSlot] = head
		}
	}
	this.bucketArr = newBucket
}

func getSlot(key, capacity int) int {
	return int(hash(key)) % capacity
}

func hash(key int) uint {
	bytes := []byte(strconv.Itoa(key))
	h := uint(5381)
	for _, b := range bytes {
		h = (h << 5) + h + uint(b)
	}
	return h
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
