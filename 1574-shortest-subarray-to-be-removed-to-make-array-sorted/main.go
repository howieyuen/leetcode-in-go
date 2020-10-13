package _574_shortest_subarray_to_be_removed_to_make_array_sorted

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}

	// 闭区间[0,left] 非递减，长度left+1
	left := 0
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}

	// 非递减序列，直接返回
	if left == n-1 {
		return 0
	}

	// 闭区间[right, n-1] 非递减, 长度n-right
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}


	// 只保留一边，计算最小删除长度
	ans := min(n-left-1, right)

	// 尝试合并，两侧均从小到大遍历
	i, j := 0, right
	for i <= left && j <= n-1 {
		if arr[i] <= arr[j] {
			ans = min(ans, j-i-1)
			i++
		} else {
			j++
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
