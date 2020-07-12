package _315_count_of_smaller_numbers_after_self

func countSmaller(nums []int) []int {
	var n = len(nums)
	if n == 0 {
		return nil
	}
	var ans = make([]int, n)
	var root = &TreeNode{Val: nums[n-1]}
	for i := n - 2; i >= 0; i-- {
		ans[i] = insert(root, nums[i])
	}
	return ans
}

type TreeNode struct {
	Val   int
	Count int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, val int) int {
	var less int
	if val <= root.Val {
		root.Count++
		if root.Left != nil {
			less += insert(root.Left, val)
		} else {
			root.Left = &TreeNode{Val: val}
		}
	} else {
		less += root.Count + 1
		if root.Right != nil {
			less += insert(root.Right, val)
		} else {
			root.Right = &TreeNode{Val: val}
		}
	}
	return less
}

func BST_insert(node *TreeNode, insert_node *TreeNode, small_count *int) {
	// 放在左边
	if insert_node.Val <= node.Val {
		node.Count++
		if node.Left != nil {
			BST_insert(node.Left, insert_node, small_count)
		} else {
			node.Left = insert_node
		}
	}
	if insert_node.Val > node.Val {
		*small_count = *small_count + node.Count + 1
		if node.Right != nil {
			BST_insert(node.Right, insert_node, small_count)
		} else {
			node.Right = insert_node
		}
	}
}

func countSmaller1(nums []int) []int {
	length := len(nums)
	count := make([]int, length)
	if length <= 1 {
		return count
	}
	
	// 从右往左处理
	node := TreeNode{Val: nums[length-1]}
	count[length-1] = 0
	
	for i := length - 2; i >= 0; i-- {
		var data int
		BST_insert(&node, &TreeNode{Val: nums[i]}, &data)
		count[i] = data
	}
	return count
}
