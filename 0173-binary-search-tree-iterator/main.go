package _173_binary_search_tree_iterator

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	values []int
	index  int
}

func inOrder(root *TreeNode) (values []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			values = append(values, root.Val)
			root = root.Right
		}
	}
	return
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		values: inOrder(root),
		index:  0,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	val := this.values[this.index]
	this.index++
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.values)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
