package _767_reorganize_string

import (
	"container/heap"
	"sort"
)

func reorganizeString(s string) string {
	n := len(s)
	var times = make([]int, 26)
	for i := range s {
		times[s[i]-'a']++
		if times[s[i]-'a'] > (n+1)/2 {
			return ""
		}
	}
	var chars = make([]byte, len(s))
	evenIdx, oddIdx := 0, 1
	for i, c := range times {
		char := byte(i + 'a')
		for c > 0 && c <= n/2 && oddIdx < n {
			chars[oddIdx] = char
			oddIdx += 2
			c--
		}
		for c > 0 {
			chars[evenIdx] = char
			c--
			evenIdx += 2
		}
	}
	return string(chars)
}

var times [26]int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return times[h.IntSlice[i]] > times[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func reorganizeString1(s string) string {
	n := len(s)
	times = [26]int{}
	for i := range s {
		times[s[i]-'a']++
		if times[s[i]-'a'] > (n+1)/2 {
			return ""
		}
	}
	h := &hp{}
	for i, c := range times {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)
	
	chars := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		chars = append(chars, byte(i+'a'), byte(j+'a'))
		if times[i]--; times[i] > 0 {
			h.push(i)
		}
		if times[j]--; times[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		chars = append(chars, byte('a'+h.IntSlice[0]))
	}
	return string(chars)
}
