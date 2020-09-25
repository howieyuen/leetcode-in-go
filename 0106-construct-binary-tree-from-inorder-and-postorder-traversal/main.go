package _106_construct_binary_tree_from_inorder_and_postorder_traversal

func buildTree(inorder []int, postorder []int) *TreeNode {
	var index = make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		index[inorder[i]] = i
	}
	return helper(index, 0, len(inorder)-1, postorder, len(postorder)-1)
}

func helper(index map[int]int, begin, end int, postOrder []int, tail int) *TreeNode {
	if begin > end {
		return nil
	}
	val := postOrder[tail]
	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	split := index[val]
	root.Left = helper(index, begin, split-1, postOrder, tail-(end-split)-1)
	root.Right = helper(index, split+1, end, postOrder, tail-1)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1
	for i := len(postorder) - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}
