package _295_find_median_from_data_stream

import "container/heap"

type MedianFinder struct {
	max *MaxHeap
	min *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	h1 := &MaxHeap{}
	heap.Init(h1)
	h2 := &MinHeap{}
	heap.Init(h2)
	return MedianFinder{
		max: h1,
		min: h2,
	}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.max, num)
	a := heap.Pop(m.max)
	heap.Push(m.min, a)
	
	if m.min.Len() > m.max.Len() {
		heap.Push(m.max, heap.Pop(m.min))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.min.Len() == m.max.Len() {
		return float64(m.max.Peek().(int)+m.min.Peek().(int)) / 2.0
	} else {
		return float64(m.max.Peek().(int))
	}
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() interface{} {
	return (*h)[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Peek() interface{} {
	return (*h)[0]
}
