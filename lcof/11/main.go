package _1

/*
 see No.154 https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
*/
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := (i + j) / 2
		if numbers[mid] > numbers[j] { // 如果mid大于right，说明还没走到右半个数组，指针i右移
			i = mid + 1
		} else if numbers[mid] < numbers[j] { // 如果mid小于right，走到右半个数组，指针j左移
			j = mid
		} else { // 如果mid等于right，说明right不是最小值，或者不是唯一的最小值，向左边缩小范围
			j--
		}
	}
	return numbers[i]
}

func minArray1(numbers []int) int {
	left, right := 0, len(numbers)-1
	if numbers[left] < numbers[right] {
		return numbers[left]
	}
	for left < right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[left] > numbers[mid] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
