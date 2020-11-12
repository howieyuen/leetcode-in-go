package _922_sort_array_by_parity_ii

func sortArrayByParityII(A []int) []int {
	for even, odd := 0, 1; even < len(A); even += 2 {
		if A[even]%2 == 1 {
			for A[odd]%2 == 1 {
				odd += 2
			}
			A[odd], A[even] = A[even], A[odd]
		}
	}
	return A
}
