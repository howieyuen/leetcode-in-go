package _378_kth_smallest_element_in_a_sorted_matrix

import (
	"container/heap"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	nums := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			nums += i + 1
			j++
		} else {
			i--
		}
	}
	return nums >= k
}

func kthSmallest1(matrix [][]int, k int) int {
	h := new(IntMaxHeap)
	heap.Init(h)
	for i := range matrix {
		for j := range matrix[i] {
			if h.Len() < k {
				heap.Push(h, matrix[i][j])
			} else if h.Peak().(int) > matrix[i][j] {
				heap.Pop(h)
				heap.Push(h, matrix[i][j])
			}
		}
	}
	return h.Peak().(int)
}

type IntMaxHeap []int

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntMaxHeap) Peak() interface{} {
	return (*h)[0]
}
