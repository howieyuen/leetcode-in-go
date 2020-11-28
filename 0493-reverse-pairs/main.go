package _493_reverse_pairs

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	n := len(nums)
	n1 := append([]int{}, nums[:n/2]...)
	n2 := append([]int{}, nums[n/2:]...)
	cnt := reversePairs(n1) + reversePairs(n2)
	
	j := 0
	for i := range n1 {
		for j < len(n2) && n1[i] > n2[j]*2 {
			j++
		}
		cnt += j
	}
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt
}
