package _099_recover_binary_search_tree

/*
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。
*/

var x, y, pred *TreeNode

func recoverTree(root *TreeNode) {
	findTwoSwapped(root)
	x.Val, y.Val = y.Val, x.Val
}

func findTwoSwapped(root *TreeNode) {
	if root == nil {
		return
	}
	findTwoSwapped(root.Left)
	if pred != nil && root.Val < pred.Val {
		y = root
		if x == nil {
			x = pred
		} else {
			return
		}
	}
	pred = root
	findTwoSwapped(root.Right)
}

func recoverTree1(root *TreeNode) {
	var stack []*TreeNode
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

func recoverTree2(root *TreeNode) {
	var nums []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findSwapped(nums)
	recoverNode(root, 2, x, y)
}

func findSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recoverNode(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recoverNode(root.Right, count, x, y)
	recoverNode(root.Left, count, x, y)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
