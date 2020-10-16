package _977_squares_of_a_sorted_array

func sortedSquares(A []int) []int {
	var left, right = 0, 0
	var ans = make([]int, len(A))
	first, last := A[0], A[len(A)-1]

	// 全>=0
	if first >= 0 {
		for i := range A {
			ans[i] = A[i] * A[i]
		}
		return ans
	}
	// 全<=0
	if last <= 0 {
		for i, j := len(A)-1, 0; i >= 0; i, j = i-1, j+1 {
			ans[j] = A[i] * A[i]
		}
		return ans
	}
	// {负,正}
	for i := 1; i < len(A); i++ {
		if first*A[i] <= 0 {
			right = i
			break
		}
	}

	i := 0
	left = right - 1
	for left >= 0 && right < len(A) {
		l, r := A[left]*A[left], A[right]*A[right]
		if l < r {
			ans[i] = l
			left--
		} else {
			ans[i] = r
			right++
		}
		i++
	}
	for left >= 0 {
		ans[i] = A[left] * A[left]
		left--
		i++
	}
	for right < len(A) {
		ans[i] = A[right] * A[right]
		right++
		i++
	}
	return ans
}

func sortedSquares1(A []int) []int {
	n := len(A)
	i, j := 0, n-1
	ans := make([]int, n)
	for index := n - 1; index >= 0; index-- {
		if left, right := A[i]*A[i], A[j]*A[j]; left > right {
			ans[index] = left
			i++
		} else {
			ans[index] = right
			j--
		}
	}
	return ans
}
