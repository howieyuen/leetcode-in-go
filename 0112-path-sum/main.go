package _112_path_sum

import (
	"container/list"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	nodeStack := list.New()
	sumStack := list.New()
	nodeStack.PushBack(root)
	sumStack.PushBack(sum)
	var curNode *TreeNode
	var curSum int
	for nodeStack.Len() > 0 {
		curNode = nodeStack.Remove(nodeStack.Front()).(*TreeNode)
		curSum = sumStack.Remove(sumStack.Front()).(int)
		if curNode.Left == nil && curNode.Right == nil && curSum == 0 {
			return true
		}
		if curNode.Left != nil {
			nodeStack.PushBack(curNode.Left)
			sumStack.PushBack(curSum - curNode.Left.Val)
		}
		if curNode.Right != nil {
			nodeStack.PushBack(curNode.Right)
			sumStack.PushBack(curSum - curNode.Right.Val)
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
