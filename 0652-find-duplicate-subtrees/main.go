package _652_find_duplicate_subtrees

import (
	"fmt"
)

var trees map[string]int
var counts map[int]int
var id int
var ans []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	trees = make(map[string]int)
	counts = make(map[int]int)
	id = 1
	ans = []*TreeNode{}
	lookUp(root)
	return ans
}

func lookUp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	key := fmt.Sprintf("%d,%d,%d", root.Val, lookUp(root.Left), lookUp(root.Right))
	var uid int
	if val, ok := trees[key]; ok {
		uid = val
	} else {
		trees[key] = id
		uid = id
		id++
	}
	if count, ok := counts[uid]; ok && count == 1 {
		ans = append(ans, root)
		counts[uid]++
	} else if !ok {
		counts[uid] = 1
	}
	return uid
}

func findDuplicateSubtrees1(root *TreeNode) []*TreeNode {
	trees = make(map[string]int)
	ans = []*TreeNode{}
	dfs(root)
	return ans
}

func dfs(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	key := fmt.Sprintf("%d,%s,%s", root.Val, dfs(root.Left), dfs(root.Right))
	if count, ok := trees[key]; ok && count == 1 {
		ans = append(ans, root)
		trees[key]++
	} else if !ok {
		trees[key] = 1
	}
	return key
}

func preOrder(root *TreeNode) []int {
	var order []int
	var stack []*TreeNode
	stack = append(stack, root)
	index := 0
	for index >= 0 {
		root = stack[index]
		order = append(order, root.Val)
		stack = stack[:index]
		index--
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return order
}

func inOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var order []int
	for root != nil || len(stack) >= 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			order = append(order, root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return order
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
