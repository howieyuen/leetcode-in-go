package _501_find_mode_in_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var repeatCounts = make(map[int]int)
	var inOrder func(root *TreeNode)
	maxTimes := -1
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		repeatCounts[root.Val]++
		maxTimes = max(maxTimes, repeatCounts[root.Val])
		inOrder(root.Right)
	}
	inOrder(root)
	var ans []int
	for val, time := range repeatCounts {
		if time == maxTimes {
			ans = append(ans, val)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
