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

func permute1(nums []int) [][]int {
	var res [][]int
	var backtrace func(nums []int, visited []bool, item []int)
	backtrace = func(nums []int, visited []bool, item []int) {
		if len(item) == len(nums) {
			var tmp = make([]int, len(item))
			copy(tmp, item)
			res = append(res, tmp)
			return
		}
		for i := range nums {
			if visited[i] {
				continue
			}
			visited[i] = true
			item = append(item, nums[i])
			backtrace(nums, visited, item)
			item = item[:len(item)-1]
			visited[i] = false
		}
	}
	var visited = make([]bool, len(nums))
	backtrace(nums, visited, nil)
	return res
}