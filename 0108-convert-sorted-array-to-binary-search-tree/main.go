package _108_convert_sorted_array_to_binary_search_tree

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedSliceToBST(nums, 0, len(nums)-1)
}

func sortedSliceToBST(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}
	mid := (i + j) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedSliceToBST(nums, i, mid-1)
	root.Right = sortedSliceToBST(nums, mid+1, j)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
