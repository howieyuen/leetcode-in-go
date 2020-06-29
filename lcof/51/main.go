package _1

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	return merge(nums, 0, len(nums)-1, make([]int, len(nums)))
}

func merge(nums []int, start, end int, tmp []int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	var count int
	count += merge(nums, start, mid, tmp)
	count += merge(nums, mid+1, end, tmp)
	index := start
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			count += mid - i + 1
			tmp[index] = nums[j]
			j++
		} else {
			tmp[index] = nums[i]
			i++
		}
		index++
	}
	for i <= mid {
		tmp[index] = nums[i]
		i++
		index++
	}
	for j <= end {
		tmp[index] = nums[j]
		j++
		index++
	}
	for k := start; k <= end; k++ {
		nums[k] = tmp[k]
	}
	return count
}
