package _700_search_in_a_binary_search_tree

/*
	二叉搜索树查找
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
