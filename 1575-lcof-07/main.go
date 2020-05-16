package _575_lcof_07

func buildTree(preorder []int, inorder []int) *TreeNode {
	splits := make(map[int]int, len(inorder))
	for i := range inorder {
		splits[inorder[i]] = i
	}
	return helper(preorder, 0, splits, 0, len(inorder)-1)
}

func helper(preOrder []int, index int, splits map[int]int, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}
	val := preOrder[index]
	split := splits[val]
	root := &TreeNode{Val: val}
	root.Left = helper(preOrder, index+1, splits, begin, split-1)
	root.Right = helper(preOrder, index+(split-begin+1), splits, split+1, end)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
