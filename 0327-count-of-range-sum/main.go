package _327_count_of_range_sum

func countRangeSum(nums []int, lower int, upper int) int {
	var mergeCount func([]int) int
	mergeCount = func(nums []int) int {
		n := len(nums)
		if n <= 1 {
			return 0
		}
		var leftHalf, rightHalf []int
		leftHalf = append(leftHalf, nums[:n/2]...)
		rightHalf = append(rightHalf, nums[n/2:]...)
		count := mergeCount(leftHalf) + mergeCount(rightHalf)
		
		var l, r int
		for _, v := range leftHalf {
			for l < len(rightHalf) && rightHalf[l]-v < lower {
				l++
			}
			for r < len(rightHalf) && rightHalf[r]-v <= upper {
				r++
			}
			count += r - l
		}
		
		// merge sort
		var p1, p2 int
		for i := range nums {
			if p1 < len(leftHalf) && (p2 == len(rightHalf) || leftHalf[p1] <= rightHalf[p2]) {
				nums[i] = leftHalf[p1]
				p1++
			} else {
				nums[i] = rightHalf[p2]
				p2++
			}
		}
		return count
	}
	
	var prefixSum = make([]int, len(nums)+1)
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return mergeCount(prefixSum)
}

func countRangeSum1(nums []int, lower int, upper int) int {
	n := len(nums)
	var ans int
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				ans++
			}
		}
	}
	return ans
}
