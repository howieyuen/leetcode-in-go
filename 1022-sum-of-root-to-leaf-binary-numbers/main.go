package _022_sum_of_root_to_leaf_binary_numbers

func sumRootToLeaf(root *TreeNode) int {
	var ans int
	dfs(root, 0, &ans)
	return ans
}

func dfs(root *TreeNode, cur int, res *int) {
	cur = (cur<<1 + root.Val) % (1e9 + 7)
	if root.Left == nil && root.Right == nil {
		*res = (*res + cur) % (1e9 + 7)
	}
	if root.Left != nil {
		dfs(root.Left, cur, res)
	}
	if root.Right != nil {
		dfs(root.Right, cur, res)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
