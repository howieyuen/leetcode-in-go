package _706_design_hashmap

type Node struct {
	Key  int
	Val  int
	Next *Node
}

const defaultCapacity = 512

type MyHashMap struct {
	nodes []*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{nodes: make([]*Node, defaultCapacity)}
}

func (this *MyHashMap) hash(key int) int {
	return key % defaultCapacity
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	node := &Node{
		Key:  key,
		Val:  value,
		Next: nil,
	}
	hash := this.hash(key)
	if this.nodes[hash] == nil {
		this.nodes[hash] = node
	} else {
		head := this.nodes[hash]
		prev := head
		for head != nil {
			if head.Key == key {
				head.Val = value
				return
			}
			prev = head
			head = head.Next
		}
		prev.Next = node
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hash := this.hash(key)
	if this.nodes[hash] == nil {
		return -1
	}
	head := this.nodes[hash]
	for head != nil {
		if head.Key == key {
			return head.Val
		}
		head = head.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hash := this.hash(key)
	if this.nodes[hash] == nil {
		return
	}
	head := this.nodes[hash]
	if head.Key == key {
		this.nodes[hash] = head.Next
		head.Next = nil
	}
	prev := head
	head = head.Next
	for head != nil {
		if head.Key == key {
			prev.Next = head.Next
			head.Next = nil
			break
		}
		prev = head
		head = head.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
