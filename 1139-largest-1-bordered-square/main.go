package _139_largest_1_bordered_square

// horizontal[i][j]表示正方形横着的边长
// vertical[i][j]表示正方形的竖着的边长
// L = min(horizontal[i][j], vertical[i][j])，可能的边长最大值是两个边长较小的那个
// 检查上边长horizontal[i-L+1][j]和左边长vertical[i][j-L+1]是否均>=L
func largest1BorderedSquare(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	var horizontal = make([][]int, rows+1)
	var vertical = make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		horizontal[i] = make([]int, cols+1)
		vertical[i] = make([]int, cols+1)
	}
	var res int
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if grid[i-1][j-1] == 0 {
				continue
			}
			horizontal[i][j] = horizontal[i][j-1] + 1
			vertical[i][j] = vertical[i-1][j] + 1
			l := min(horizontal[i][j], vertical[i][j])
			for l > 0 {
				if horizontal[i-l+1][j] >= l && vertical[i][j-l+1] >= l {
					break
				}
				l--
			}
			res = max(res, l)
		}
	}
	return res * res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
