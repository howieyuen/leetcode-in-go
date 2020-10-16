package _442_count_triplets_that_can_form_two_arrays_of_equal_xor

// 0 <= i < j <= k < arr.length
// a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
// b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
// 要求 a == b，即 a ^ b = 0，[i, k]范围内的所有元素的异或结果为0，
// 那么j的位置在(i,k]之间任选，共有k-i个分割点
func countTriplets(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}
	var ans int
	for i := 0; i < n-1; i++ {
		cur := 0
		for k := i; k < n; k++ {
			cur ^= arr[k]
			if cur == 0 && k > i {
				ans += k - i
			}
		}
	}
	return ans
}
