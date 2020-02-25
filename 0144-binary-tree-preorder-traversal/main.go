package _144_binary_tree_preorder_traversal

/*
	给定一个二叉树，返回它的 前序 遍历。
*/
func preorderTraversal(root *TreeNode) []int {
	var order []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			order = append(order, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return order
}

func preorderTraversal2(root *TreeNode) []int {
	var ans []int
	if root != nil {
		stack := []*TreeNode{}
		stack = append(stack, root)
		for len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return ans
}

func preorderTraversal1(root *TreeNode) []int {
	var ans []int
	preOrder(root, &ans)
	return ans
}

func preOrder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}
	*order = append(*order, root.Val)
	if root.Left != nil {
		preOrder(root.Left, order)
	}
	if root.Right != nil {
		preOrder(root.Right, order)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
