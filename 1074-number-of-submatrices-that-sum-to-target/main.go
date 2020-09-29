package _074_number_of_submatrices_that_sum_to_target

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	rows, cols := len(matrix), len(matrix[0])
	var sums = make([][]int, rows+1)
	for i := range sums {
		sums[i] = make([]int, cols+1)
	}
	// 计算前缀和
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] + matrix[i-1][j-1] - sums[i-1][j-1]
		}
	}
	var ans int
	// (  ,   ) (  ,   ) ... (  ,   ) (  ,   )
	// (  ,   ) (x1, y1) ... (x1, y2) (  ,   )
	// (  ,   ) (  ,   ) ... (  ,   ) (  ,   )
	// (  ,   ) (x2, y1) ... (x2, y2) (  ,   )
	// (  ,   ) (  ,   ) ... (  ,   ) (  ,   )
	for x1 := 0; x1 < rows; x1++ { // 左上角
		for y1 := 0; y1 < cols; y1++ {
			for x2 := x1; x2 < rows; x2++ { // 右下角
				for y2 := y1; y2 < cols; y2++ {
					if sums[x2+1][y2+1]-sums[x2+1][y1]-sums[x1][y2+1]+sums[x1][y1] == target {
						ans++
					}
				}
			}
		}
	}
	return ans
}
