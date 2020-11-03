package _941_valid_mountain_array

func validMountainArray(A []int) bool {
	left, right := 0, len(A)-1
	for left < right && A[left] < A[left+1] {
		left++
	}
	for left < right && A[right-1] > A[right] {
		right--
	}
	return left == right && left != 0 && right != len(A)-1
}
