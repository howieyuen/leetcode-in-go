package _215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	buildHeap(nums)
	n := len(nums)
	for k > 1 {
		nums[n-1], nums[0] = nums[0], nums[n-1]
		k--
		n--
		heapify(nums[:n], 0)
	}
	return nums[0]
}

func buildHeap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i)
	}
}

// 最大堆
func heapify(nums []int, i int) {
	left, right := i*2+1, i*2+2
	largest := i
	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, largest)
	}
}
