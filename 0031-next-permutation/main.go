package _031_next_permutation

/*
从后往前扫描找到第一个下标i,使得nums[i] < nums[i + 1];
若i = -1，表示该数组为单调递减序列;否则，从后往前扫描找到找到第一个下标j使得nums[j] > nums[i]，交换两数且不破坏nums(i, nums.length)的单调性；
翻转nums(i, nums.length)使其成为单调增序列。
*/
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	for a, b := start, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}
