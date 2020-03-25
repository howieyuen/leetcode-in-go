package _059_spiral_matrix_ii

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	num := 1
	left, up := 0, 0
	right, down := n-1, n-1
	for num <= n*n {
		for j := left; j <= right; j++ {
			ans[up][j] = num
			num++
		}
		up++
		for i := up; i <= down; i++ {
			ans[i][right] = num
			num++
		}
		right--
		for j := right; j >= left; j-- {
			ans[down][j] = num
			num++
		}
		down--
		for i := down; i >= up; i-- {
			ans[i][left] = num
			num++
		}
		left++
	}
	return ans
}
