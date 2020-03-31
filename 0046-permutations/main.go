package _046_permutations

func permute(nums []int) [][]int {
	var ans [][]int
	var backtrace func(first int, n int, cur []int)
	backtrace = func(first int, n int, cur []int) {
		if first == n {
			var tmp = make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
		}
		for i := first; i < n; i++ {
			cur[i], cur[first] = cur[first], cur[i]
			backtrace(first+1, n, cur)
			cur[i], cur[first] = cur[first], cur[i]
		}
	}
	cur := make([]int, len(nums))
	copy(cur, nums)
	backtrace(0, len(nums), cur)
	return ans
}
