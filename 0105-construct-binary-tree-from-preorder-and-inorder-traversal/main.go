package _105_construct_binary_tree_from_preorder_and_inorder_traversal

func buildTree(preorder []int, inorder []int) *TreeNode {
	var indexes = make(map[int]int)
	for i := range inorder {
		indexes[inorder[i]] = i
	}
	return helper(indexes, 0, len(indexes)-1, preorder, 0)
}

func helper(indexes map[int]int, begin, end int, preOrder []int, front int) *TreeNode {
	if begin > end {
		return nil
	}
	val := preOrder[front]
	split := indexes[val]
	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	root.Left = helper(indexes, begin, split-1, preOrder, front+1)
	root.Right = helper(indexes, split+1, end, preOrder, front+(split-begin)+1)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
