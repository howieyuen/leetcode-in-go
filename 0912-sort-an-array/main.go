package _912_sort_an_array

// quick sort
func sortArray1(nums []int) []int {
	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		i, j := left, right
		pivot := nums[left]
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			for i < j && nums[i] < pivot {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
		}
		if i > left {
			quickSort(left, i-1)
		}
		if j < right {
			quickSort(j+1, right)
		}
	}
	quickSort(0, len(nums)-1)
	return nums
}

// merge sort
func sortArray2(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, low, high int) {
	mid := (low + high) / 2
	if low < high {
		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
}

func merge(nums []int, low, mid, high int) {
	tmp := make([]int, high-low+1)
	i, j, index := low, mid+1, 0
	for i <= mid && j <= high {
		if nums[i] < nums[j] {
			tmp[index] = nums[i]
			i++
		} else {
			tmp[index] = nums[j]
			j++
		}
		index++
	}
	for i <= mid {
		tmp[index] = nums[i]
		i++
		index++
	}
	for j <= high {
		tmp[index] = nums[j]
		j++
		index++
	}
	if len(tmp) > 0 {
		copy(nums[low:high+1], tmp)
	}
}

// heap sort
func sortArray3(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	// 从最后一个非叶节点开始向前遍历，调整节点性质，使之成为大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
	// 交换堆顶和当前末尾的节点，重置大顶堆
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		heapify(nums, 0, n)
	}
}

func heapify(nums []int, i, n int) {
	// 先根据堆性质，找出它左右节点的索引
	left := 2*i + 1
	right := 2*i + 2
	// 默认当前节点（父节点）是最大值。
	largestIndex := i
	if left < n && nums[left] > nums[largestIndex] {
		// 如果有左节点，并且左节点的值更大，更新最大值的索引
		largestIndex = left
	}
	if right < n && nums[right] > nums[largestIndex] {
		// 如果有右节点，并且右节点的值更大，更新最大值的索引
		largestIndex = right
	}
	if largestIndex != i {
		// 如果最大值不是当前非叶子节点的值，那么就把当前节点和最大值的子节点值互换
		nums[largestIndex], nums[i] = nums[i], nums[largestIndex]
		// 因为互换之后，子节点的值变了，如果该子节点也有自己的子节点，仍需要再次调整。
		heapify(nums, largestIndex, n)
	}
}

// bubble sort
func sortArray4(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// insert sort
func sortArray5(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		var j int
		for j = i; j > 0 && cur < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = cur
	}
	return nums
}

// select sort
func sortArray6(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		k := i
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	return nums
}

// shell sort
func sortArray7(nums []int) []int {
	for increment := len(nums) / 2; increment > 0; increment /= 2 {
		for i := increment; i < len(nums); i++ {
			tmp := nums[i]
			for j := i; j >= increment; j -= increment {
				if tmp < nums[j-increment] {
					nums[j] = nums[j-increment]
				} else {
					nums[j] = tmp
					break
				}
			}
		}
	}
	return nums
}
