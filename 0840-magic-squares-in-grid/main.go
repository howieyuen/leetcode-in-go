package _840_magic_squares_in_grid

func numMagicSquaresInside(grid [][]int) int {
	var check = func(i, j int) bool {
		oblique := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
		if oblique != grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] {
			return false
		}
		var visited = make([]bool, 10)
		for deltaX := 0; deltaX < 3; deltaX++ {
			row, col := 0, 0
			for deltaY := 0; deltaY < 3; deltaY++ {
				num := grid[i+deltaX][j+deltaY]
				if num > 9 || num < 1 || visited[num] {
					return false
				}
				visited[num] = true
				row += num
				col += grid[i+deltaY][j+deltaX]
			}
			if row != oblique || col != oblique {
				return false
			}
		}
		return true
	}
	
	var ans int
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if check(i, j) {
				ans++
			}
		}
	}
	return ans
}
