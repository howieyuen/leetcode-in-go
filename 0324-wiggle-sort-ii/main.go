package _324_wiggle_sort_ii

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	right, mid := n, (n+1)>>1
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		if i&1 == 1 {
			right--
			nums[i] = tmp[right]
		} else {
			mid--
			nums[i] = tmp[mid]
		}
	}
}

func wiggleSort1(nums []int) {
	n := len(nums)
	quickSelect(nums, 0, n, n/2)
	mid := nums[n/2]
	i, j, k := 0, 0, n-1
	for j <= k {
		if nums[j] > mid {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else if nums[j] < mid {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	var tmp1 = append([]int{}, nums[:(n+1)/2]...)
	var tmp2 = append([]int{}, nums[(n+1)/2:]...)
	for i := 0; i < len(tmp1); i++ {
		nums[2*i] = tmp1[len(tmp1)-i-1]
	}
	for i := 0; i < len(tmp2); i++ {
		nums[2*i+1] = tmp2[len(tmp2)-i-1]
	}
}

func quickSelect(nums []int, begin, end int, n int) {
	t := nums[end-1]
	i, j := begin, begin
	for j < end {
		if nums[j] <= t {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	if i-1 > n {
		quickSelect(nums, begin, i-1, n)
	} else if i <= n {
		quickSelect(nums, i, end, n)
	}
}