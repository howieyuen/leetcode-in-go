package _275_h_index_ii

// 给定一个大小为 n 的升序的引用次数列表
// 找到满足 citations[i] >= n - i 的第一个 citations[i]
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] == n-mid {
			return n - mid
		} else if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}
