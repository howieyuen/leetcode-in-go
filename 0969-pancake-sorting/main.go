package _969_pancake_sorting

func pancakeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return nil
	}
	
	var ans []int
	for j := 0; j < n; j++ {
		k, max := 0, arr[0]
		for i := 1; i < n-j; i++ {
			if max < arr[i] {
				k = i
				max = arr[i]
			}
		}
		if k == n-j-1 {
			continue
		}
		if k != 0 {
			ans = append(ans, k+1)
			reverse(arr, k)
		}
		ans = append(ans, n-j)
		reverse(arr, n-j-1)
	}
	return ans
}

func reverse(arr []int, k int) {
	for i, j := 0, k; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
