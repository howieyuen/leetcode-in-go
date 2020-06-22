package _687_longest_univalue_path

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var longest = 1
	var bfs func(root *TreeNode, curVal int) int
	// 当root为空，bfs返回0
	bfs = func(root *TreeNode, curVal int) int {
		if root == nil {
			return 0
		}
		left := bfs(root.Left, root.Val)
		right := bfs(root.Right, root.Val)
		// 路径长度为节点数减1所以此处不加1
		longest = max(longest, left+right)
		if root.Val == curVal {
			return max(left, right) + 1
		}
		return 0
	}
	bfs(root, root.Val)
	return longest
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
