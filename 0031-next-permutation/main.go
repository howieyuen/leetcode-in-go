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

func nextPermutation1(nums []int) {
	i := len(nums) - 2
	// 逆序查找第一个比后者小的元素nums[i]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		// 序列非降序，从后向前找到第一个比nums[i]大的元素nums[j]
		j := len(nums) - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 交换两者，因为nums[j]是第一个比nums[j]大的元素，而nums[j:n]之间是降序，交换之后，仍然是降序
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 反转nums[i+1:n]，即可保证右半部分数组为升序
	for a, b := i+1, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}
