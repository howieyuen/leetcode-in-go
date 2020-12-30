package _046_last_stone_weight

import (
	"container/heap"
	"sort"
)

func lastStoneWeight(stones []int) int {
	q := &IntHeap{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}

type IntHeap struct {
	sort.IntSlice
}

func (h IntHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}

func (h *IntHeap) push(x int) {
	heap.Push(h, x)
}

func (h *IntHeap) pop() int {
	return heap.Pop(h).(int)
}
