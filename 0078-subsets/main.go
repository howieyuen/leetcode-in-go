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
