package _146_lru_cache

/*
LRU设计，要求Get和Put均为O(1)

维护map，实现Get的O(1)
维护双链表，实现Put的O(1)

新插入和Get的节点，总是置为head
表满时，总时删除tail
*/

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:     nil,
		tail:     nil,
		cache:    make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.renewHead(node)
	return node.val
}

func (this *LRUCache) remove(node *Node) {
	if node.pre == nil {
		this.head = node.next
	} else {
		node.pre.next = node.next
	}
	if node.next == nil {
		this.tail = node.pre
	} else {
		node.next.pre = node.pre
	}
	node.next = nil
	node.pre = nil
}

func (this *LRUCache) renewHead(node *Node) {
	if this.head == nil {
		this.head = node
	} else {
		this.head.pre = node
		node.next = this.head
		this.head = node
	}
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.remove(node)
		this.renewHead(node)
	} else {
		if len(this.cache) == this.capacity {
			delete(this.cache, this.tail.key)
			this.remove(this.tail)
		}
		node := &Node{key: key, val: value}
		this.cache[key] = node
		this.renewHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
