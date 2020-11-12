package _513_find_bottom_left_tree_value

func findBottomLeftValue(root *TreeNode) int {
	var maxDepth = -1
	var leftest int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if depth > maxDepth {
				leftest = root.Val
				maxDepth = depth
			}
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return leftest
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
