package _119_pascals_triangle_ii

func getRow(rowIndex int) []int {
	ans := []int{1}
	for i := 1; i <= rowIndex; i++ {
		ans = append(ans, 1)
		for j := i - 1; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}
	return ans
}

/*
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1



*/
