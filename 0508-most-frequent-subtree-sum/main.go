package _508_most_frequent_subtree_sum

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var times = make(map[int]int)
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		var left, right int
		if root.Left != nil {
			left = dfs(root.Left)
		}
		if root.Right != nil {
			right = dfs(root.Right)
		}
		times[left+right+root.Val]++
		return left + right + root.Val
	}
	dfs(root)
	var ans []int
	var maxTime = 1
	for key, time := range times {
		if time == maxTime {
			ans = append(ans, key)
		} else if time > maxTime {
			maxTime = time
			ans = []int{key}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
