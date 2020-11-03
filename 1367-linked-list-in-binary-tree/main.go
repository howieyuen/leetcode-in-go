package _367_linked_list_in_binary_tree

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(head *ListNode, root *TreeNode) bool
	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil {
			return false
		}
		if root.Val != head.Val {
			return false
		}
		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}
	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
