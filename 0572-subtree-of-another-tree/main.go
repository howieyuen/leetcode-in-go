package _572_subtree_of_another_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	var checkFunc func(s *TreeNode, t *TreeNode) bool
	checkFunc = func(s *TreeNode, t *TreeNode) bool {
		if s == nil && t == nil {
			return true
		}
		if s == nil || t == nil {
			return false
		}
		if s.Val != t.Val {
			return false
		}
		return checkFunc(s.Left, t.Left) && checkFunc(s.Right, t.Right)
	}
	return checkFunc(s, t)||isSubtree(s.Left, t)||isSubtree(s.Right, t)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
