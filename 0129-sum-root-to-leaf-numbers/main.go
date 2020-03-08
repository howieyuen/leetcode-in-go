package _129_sum_root_to_leaf_numbers

/*
求根到叶子节点数字之和
*/
func sumNumbers(root *TreeNode) int {
	return bfs(root, 0)
}

func bfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return bfs(root.Left, sum) + bfs(root.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
