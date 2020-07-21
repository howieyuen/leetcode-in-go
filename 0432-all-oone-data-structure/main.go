package _432_all_oone_data_structure

type Node struct {
	Val  int
	prev *Node
	next *Node
}

type DLinkedList struct {
	head *Node
	tail *Node
}

func (dl *DLinkedList) PushFront(n *Node) {
	n.prev = nil
	if dl.head != nil {
		dl.head.prev = n
	} else {
		dl.tail = n
	}
	n.next = dl.head
	dl.head = n
}

func (dl *DLinkedList) InsertBefore(n, before *Node) {
	if before != dl.head {
		n.prev, n.next = before.prev, before
		n.prev.next, before.prev = n, n
	} else {
		n.prev, n.next = nil, before
		before.prev, dl.head = n, n
	}
}

func (dl *DLinkedList) InsertAfter(n, after *Node) {
	if after != dl.tail {
		n.prev, n.next = after, after.next
		n.next.prev, after.next = n, n
	} else {
		n.prev, n.next = after, nil
		after.next, dl.tail = n, n
	}
}

func (dl *DLinkedList) PopFront() *Node {
	if dl.head == nil {
		return nil
	}
	h := dl.head
	if h.next != nil {
		dl.head = h.next
	} else {
		dl.head, dl.tail = nil, nil
	}
	return h
}

func (dl *DLinkedList) PopBack() *Node {
	if dl.tail == nil {
		return nil
	}
	r := dl.tail
	if r.prev != nil {
		dl.tail = r.prev
	} else {
		dl.head, dl.tail = nil, nil
	}
	return r
}

func (dl *DLinkedList) Unlink(n *Node) {
	if dl.head == n {
		dl.PopFront()
		return
	} else if dl.tail == n {
		dl.PopBack()
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

type AllOne struct {
	keyToVal    map[string]int
	valToKeys   map[int]map[string]bool
	keyToNode   map[int]*Node
	dLinkedList *DLinkedList
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		keyToVal:    map[string]int{},
		valToKeys:   map[int]map[string]bool{},
		keyToNode:   map[int]*Node{},
		dLinkedList: &DLinkedList{},
	}
}

func (o *AllOne) addKV(key string, val int, onInc bool) {
	if m, ok := o.valToKeys[val]; ok {
		m[key] = true
	} else {
		o.valToKeys[val] = map[string]bool{key: true}
	}
	
	if len(o.valToKeys[val]) == 1 {
		n := &Node{Val: val}
		o.keyToNode[val] = n
		if onInc {
			if val != 1 {
				o.dLinkedList.InsertAfter(n, o.keyToNode[val-1])
			} else {
				o.dLinkedList.PushFront(n)
			}
		} else {
			o.dLinkedList.InsertBefore(n, o.keyToNode[val+1])
		}
	}
}

func (o *AllOne) removeKV(key string, val int) {
	m := o.valToKeys[val]
	if len(m) != 1 {
		delete(m, key)
	} else {
		delete(o.valToKeys, val)
		o.dLinkedList.Unlink(o.keyToNode[val])
		delete(o.keyToNode, val)
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (o *AllOne) Inc(key string) {
	o.keyToVal[key]++
	v := o.keyToVal[key]
	o.addKV(key, v, true)
	if v != 1 {
		o.removeKV(key, v-1)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (o *AllOne) Dec(key string) {
	if _, ok := o.keyToVal[key]; !ok {
		return
	}
	o.keyToVal[key]--
	v := o.keyToVal[key]
	if v != 0 {
		o.addKV(key, v, false)
	}
	o.removeKV(key, v+1)
}

func (o *AllOne) GetOneKey(n *Node) string {
	if n == nil {
		return ""
	}
	keys := o.valToKeys[n.Val]
	for k := range keys {
		return k
	}
	return ""
}

/** Returns one of the keys with maximal value. */
func (o *AllOne) GetMaxKey() string {
	return o.GetOneKey(o.dLinkedList.tail)
}

/** Returns one of the keys with Minimal value. */
func (o *AllOne) GetMinKey() string {
	return o.GetOneKey(o.dLinkedList.head)
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
