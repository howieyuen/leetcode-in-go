package _114_flatten_binary_tree_to_linked_list

// 后续遍历 迭代
func flatten2(root *TreeNode) {
	var stack []*TreeNode
	var cur = root
	var pre *TreeNode = nil
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		if cur.Left == nil || cur.Left == pre {
			stack = stack[:len(stack)-1]
			cur.Right = pre
			cur.Left = nil
			pre = cur
			cur = nil
		} else {
			cur = cur.Left
		}
	}
}

// 后序遍历 递归
func flatten1(root *TreeNode) {
	var pre *TreeNode = nil
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Right)
		postOrder(node.Left)
		node.Right = pre
		node.Left = nil
		pre = node
	}
	postOrder(root)
}

func flatten(root *TreeNode) {
	for root != nil {
		// 左子树为 null，直接考虑下一个节点
		if root.Left == nil {
			root = root.Right
		} else {
			var pre = root.Left
			// 找左子树最右边的节点
			for pre.Right != nil {
				pre = pre.Right
			}
			// 将原来的右子树接到左子树的最右边节点
			pre.Right = root.Right
			// 将左子树插入到右子树的地方
			root.Right = root.Left
			// 左子树置为空
			root.Left = nil
			// 考虑下一个节点
			root = root.Right
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
