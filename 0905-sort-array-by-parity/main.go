package _905_sort_array_by_parity

func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		for i < j && A[i]%2 == 0 {
			i++
		}
		for i < j && A[j]%2 == 1 {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
