package _632_smallest_range_covering_elements_from_k_lists

import (
	"container/heap"
	"math"
)

/*
给k个升序排序的数组，找到一个最小区间，使得每个数组中至少有一个数在此区间内

问题转化为：k个数组中，各取一个数，使得这组数的最大值和最小值的差最小
*/

var (
	// numsC 用于最小堆调整堆顶元素
	numsC [][]int
	// next 表示数组的元素的具体下标，即二维数组的列号
	next  []int
)

func smallestRange(nums [][]int) []int {
	numsC = nums
	//
	left, right := math.MinInt32, math.MaxInt32
	minRange := right - left
	
	max := math.MinInt32
	next = make([]int, len(nums))
	
	// 最小堆，取nums[i][next[i]]和nums[j][next[j]比大小得出
	h := &IntHeap{}
	
	// 每个数组的首位置指针入堆，即二维数组的行号
	for i := range nums {
		heap.Push(h, i)
		max = Max(max, nums[i][0])
	}
	
	for {
		// 取出堆顶元素
		minIndex := heap.Pop(h).(int)
		curRange := max - nums[minIndex][next[minIndex]]
		// 范围缩小
		if curRange < minRange {
			minRange = curRange
			left, right = nums[minIndex][next[minIndex]], max
		}
		// 堆顶元素所在的数组指针后移
		next[minIndex]++
		// 已有数组遍历结束，返回结果
		if next[minIndex] == len(nums[minIndex]) {
			break
		}
		// 指针后移后指向的元素入堆
		heap.Push(h, minIndex)
		// 更新最大值
		max = Max(max, nums[minIndex][next[minIndex]])
	}
	return []int{left, right}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
