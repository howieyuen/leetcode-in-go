package _054_distant_barcodes

import (
	"container/heap"
	"sort"
)

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	times := make(map[int]int)
	for i := range barcodes {
		times[barcodes[i]]++
	}
	var ans = make([]int, n)
	evenIdx, oddIdx := 0, 1
	for k := range times {
		for times[k] > 0 && times[k] <= n/2 && oddIdx < n {
			ans[oddIdx] = k
			oddIdx += 2
			times[k]--
		}
		for times[k] > 0 {
			ans[evenIdx] = k
			times[k]--
			evenIdx += 2
		}
	}
	return ans
}

func rearrangeBarcodes1(barcodes []int) []int {
	n := len(barcodes)
	counts = make(map[int]int)
	for i := range barcodes {
		counts[barcodes[i]]++
	}
	h := &hp{}
	for key := range counts {
		h.IntSlice = append(h.IntSlice, key)
	}
	heap.Init(h)
	var ans = make([]int, 0, n)
	for len(h.IntSlice) > 1 {
		key1, key2 := h.pop(), h.pop()
		ans = append(ans, key1, key2)
		if counts[key1]--; counts[key1] > 0 {
			h.push(key1)
		}
		if counts[key2]--; counts[key2] > 0 {
			h.push(key2)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, h.pop())
	}
	return ans
}

var counts map[int]int

type hp struct {
	sort.IntSlice
}

func (hp *hp) Less(i, j int) bool {
	return counts[hp.IntSlice[i]] > counts[hp.IntSlice[j]]
}

func (hp *hp) Push(v interface{}) {
	hp.IntSlice = append(hp.IntSlice, v.(int))
}

func (hp *hp) Pop() interface{} {
	v := hp.IntSlice[len(hp.IntSlice)-1]
	hp.IntSlice = hp.IntSlice[:len(hp.IntSlice)-1]
	return v
}

func (hp *hp) push(v int) {
	heap.Push(hp, v)
}

func (hp *hp) pop() int {
	return heap.Pop(hp).(int)
}
