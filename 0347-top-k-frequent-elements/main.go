package _347_top_k_frequent_elements

import (
	"container/heap"
	"sort"
)

func topKFrequent1(nums []int, k int) []int {
	var res []int
	maps := make(map[int]int)
	for _, n := range nums {
		maps[n]++
	}
	
	counts := make([]int, 0, len(maps))
	for _, c := range maps {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	min := counts[len(counts)-k]
	
	for n, c := range maps {
		if c >= min {
			res = append(res, n)
		}
	}
	return res
}

type Node struct {
	val  int
	freq int
}

type NodeList []*Node

func (nodeList *NodeList) Len() int {
	return len(*nodeList)
}

func (nodeList *NodeList) Swap(i, j int) {
	(*nodeList)[i], (*nodeList)[j] = (*nodeList)[j], (*nodeList)[i]
}

func (nodeList *NodeList) Less(i, j int) bool {
	return (*nodeList)[i].freq < (*nodeList)[j].freq
}

func (nodeList *NodeList) Push(node interface{}) {
	*nodeList = append(*nodeList, node.(*Node))
}

func (nodeList *NodeList) Pop() interface{} {
	node := (*nodeList)[nodeList.Len()-1]
	*nodeList = (*nodeList)[:nodeList.Len()-1]
	return node
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	freq := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	nodeList := &NodeList{}
	heap.Init(nodeList)
	
	for key, v := range freq {
		node := &Node{
			val:  key,
			freq: v,
		}
		heap.Push(nodeList, node)
		if nodeList.Len() > k {
			heap.Pop(nodeList)
		}
	}
	for nodeList.Len() > 0 {
		node := nodeList.Pop().(*Node)
		result = append(result, node.val)
	}
	return result
}
