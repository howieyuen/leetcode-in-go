package _053_maximum_subarray

func maxSubArray(nums []int) int {
	ans := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], cur+nums[i])
		ans = max(ans, cur)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Node struct {
	leftSum  int
	rightSum int
	maxSum   int
	totalSum int
}

func newStatus(l, r, m, t int) Node {
	return Node{
		leftSum:  l,
		rightSum: r,
		maxSum:   m,
		totalSum: t,
	}
}

func maxSubArray1(nums []int) int {
	return shiftUp(nums, 0, len(nums)-1).maxSum
}

func shiftUp(nums []int, left, right int) Node {
	if left == right {
		return newStatus(nums[left], nums[left], nums[left], nums[left])
	}
	mid := (left + right) >> 1
	leftSub := shiftUp(nums, left, mid)
	rightSub := shiftUp(nums, mid+1, right)
	return combineStatus(leftSub, rightSub)
}

func combineStatus(left, right Node) Node {
	leftSum := max(left.leftSum, left.totalSum+right.leftSum)
	rightSum := max(right.rightSum, right.totalSum+left.rightSum)
	maxSum := max(max(left.maxSum, right.maxSum), left.rightSum+right.leftSum)
	totalSum := left.totalSum + right.totalSum
	return newStatus(leftSum, rightSum, maxSum, totalSum)
}
