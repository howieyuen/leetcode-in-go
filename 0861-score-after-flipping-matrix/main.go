package _861_score_after_flipping_matrix

func matrixScore(A [][]int) int {
	rows, cols := len(A), len(A[0])
	for i := 0; i < rows; i++ {
		if A[i][0] == 0 {
			for j := 1; j < cols; j++ {
				A[i][j] ^= 1
			}
		}
	}
	var ans = rows << (cols - 1)
	for j := 1; j < cols; j++ {
		cnt := 0
		for i := 0; i < rows; i++ {
			if A[i][j] == 1 {
				cnt++
			}
		}
		if cnt < rows-cnt {
			cnt = rows - cnt
		}
		ans += cnt << (cols - j - 1)
	}
	return ans
}
