package test1

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var ans int
	n := len(arr)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[i]-arr[k]) <= c && abs(arr[j]-arr[k]) <= b {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
