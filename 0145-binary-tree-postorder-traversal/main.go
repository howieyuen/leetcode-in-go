package _145_binary_tree_postorder_traversal

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(ans []int) {
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
}

func postorderTraversal1(root *TreeNode) []int {
	var ans []int
	postOrder(root, &ans)
	return ans
}

func postOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		postOrder(root.Left, ans)
	}
	if root.Right != nil {
		postOrder(root.Right, ans)
	}
	*ans = append(*ans, root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
