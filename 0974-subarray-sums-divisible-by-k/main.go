package _974_subarray_sums_divisible_by_k

/*
974：设置数据量，暴力会超时
see 560
*/
func subarraysDivByK(a []int, k int) int {
	var store = make(map[int]int)
	store[0] = 1
	sum := 0
	cnt := 0
	for i := range a {
		sum += a[i]
		t := sum % k
		if t < 0 {
			t += k
		}
		cnt += store[t]
		store[t]++
	}
	return cnt
}
