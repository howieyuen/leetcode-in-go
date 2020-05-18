package _6

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}
	return isSub(a, b) || isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func isSub(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
