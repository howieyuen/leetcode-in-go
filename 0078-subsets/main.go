package _078_subsets

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, num := range nums {
		for _, sub := range res {
			clone := make([]int, len(sub))
			copy(clone, sub)
			clone = append(clone, num)
			res = append(res, clone)
		}
	}
	return res
}

func subsets1(nums []int) [][]int {
	var ans [][]int
	var backtrace func(index int, size int, cur []int)
	backtrace = func(index int, size int, cur []int) {
		if len(cur) == size {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
		}
		for i := index; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrace(i+1, size, cur)
			cur = cur[:len(cur)-1]
		}
	}
	for i := 0; i <= len(nums); i++ {
		backtrace(0, i, nil)
	}
	return ans
}
