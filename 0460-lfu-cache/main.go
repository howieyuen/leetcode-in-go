package _460_lfu_cache

type Node struct {
	key  int
	val  int
	freq int
	next *Node
	prev *Node
}

func (n *Node) insert(node *Node) {
	node.next = n
	node.prev = n.prev
	n.prev.next = node
	n.prev = node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	list := &LinkedList{
		head: &Node{},
		tail: &Node{},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keysTable: make(map[int]*Node),
		freqTable: make([]*LinkedList, capacity+1),
		capacity:  capacity,
		minFreq:   1,
	}
}

type LFUCache struct {
	keysTable map[int]*Node
	freqTable []*LinkedList
	minFreq   int
	size      int
	capacity  int
}

func (c *LFUCache) unlink(node *Node) int {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.freq > 0 && node.prev == c.freqTable[node.freq].head && node.next == c.freqTable[node.freq].tail {
		c.freqTable[node.freq] = nil
	}
	node.prev = nil
	node.next = nil
	return node.key
}

func (c *LFUCache) updateFreq(node *Node) {
	c.unlink(node)
	
	node.freq++
	// 2倍扩容
	if node.freq >= len(c.freqTable) {
		c.freqTable = append(c.freqTable, make([]*LinkedList, node.freq*2)...)
	}
	if c.freqTable[node.freq] == nil {
		c.freqTable[node.freq] = NewLinkedList()
	}
	
	c.freqTable[node.freq].tail.insert(node)
	
	if node.freq < c.minFreq || c.freqTable[c.minFreq] == nil {
		c.minFreq = node.freq
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.keysTable[key]
	if !ok {
		return -1
	}
	
	c.updateFreq(node)
	return node.val
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity < 1 {
		return
	}
	node, ok := c.keysTable[key]
	if !ok {
		node = &Node{
			key: key,
			val: value,
		}
		c.keysTable[key] = node
		c.size++
	} else {
		node.val = value
	}
	
	if c.size > c.capacity {
		delete(c.keysTable, c.unlink(c.freqTable[c.minFreq].head.next))
		c.size--
	}
	
	c.updateFreq(node)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
