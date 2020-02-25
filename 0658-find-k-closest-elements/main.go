package _658_find_k_closest_elements

/*
	二分查找到x的下标index，按照[index-k+1，index+k-1]缩小窗口，直到size=k
*/
func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	if arr[left] >= x {
		return arr[:k]
	}
	right := len(arr) - 1
	if arr[right] <= x {
		return arr[len(arr)-k:]
	}
	index := binarySearch(arr, left, right, x)
	if index < 0 {
		index = -index - 1
	}
	left, right = max(0, index-k-1), min(len(arr)-1, index+k-1)
	for right-left > k-1 {
		if x-arr[left] <= arr[right]-x {
			right--
		} else {
			left++
		}
	}
	return arr[left : right+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func binarySearch(nums []int, start, end, target int) int {
	for start <= end {
		var mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -(start + 1)
}
