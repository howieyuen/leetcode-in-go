package _047_permutations_ii

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var backtrace func(index int, n int)
	backtrace = func(index int, n int) {
		if index == n {
			var tmp = make([]int, n)
			copy(tmp, nums)
			ans = append(ans, tmp)
		}
		used := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			nums[i], nums[index] = nums[index], nums[i]
			backtrace(index+1, n)
			nums[i], nums[index] = nums[index], nums[i]
			used[nums[i]] = true
		}
	}
	backtrace(0, len(nums))
	return ans
}
