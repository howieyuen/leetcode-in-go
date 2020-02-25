package _410_split_array_largest_sum

/*
给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

1. 分成子数组，最多就分成len(nums)个，此时Max(sum) = Max(nums)，此时和最小；最少就分成一个组，max(sum) = Sum(nums) ，此时和为最大
2. 结果必然在此二者之间，二分法寻找
3. 确定mid，子数组是连续的，即遍历子数组，当前和大于mid时，数组个数group++
4. 比较数组个数和要求m的大小，如group > m, 说明分组太多，mid值太小，不需要几个数字组合即可满足，所以范围右移，反之左移
*/
func splitArray(nums []int, m int) int {
	var left, right int
	for _, num := range nums {
		right += num
		if left < num {
			left = num
		}
	}
	var mid int
	for left < right {
		mid = (left + right) >> 1
		groups := 1
		curSum := 0
		for _, num := range nums {
			curSum += num
			if curSum > mid {
				groups++
				curSum = num
			}
		}
		if groups > m { // 分组过多，需要减少分组，即group++的条件要少满足，即mid要增大，范围要向右移动
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
