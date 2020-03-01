package _450_delete_node_in_a_bst

/*
	二叉搜索树删除节点
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left != nil {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, key)
		} else {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, key)
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
