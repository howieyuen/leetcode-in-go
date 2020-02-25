package _733_flood_fill

/*
	从指定位置开始，在上下左右四个方向上直接连接的位置，将与开始位置的值相同的点，均更新成新值

	DFS 从开始位置更新新值，从四个方向上DFS，直到结束
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		dfs(image, sr, sc, color, newColor)
	}
	return image
}

func dfs(image [][]int, i, j int, color, newColor int) {
	if image[i][j] == color {
		image[i][j] = newColor
		if i > 0 {
			dfs(image, i-1, j, color, newColor)
		}
		if j > 0 {
			dfs(image, i, j-1, color, newColor)
		}
		if i < len(image)-1 {
			dfs(image, i+1, j, color, newColor)
		}
		if j < len(image[0])-1 {
			dfs(image, i, j+1, color, newColor)
		}
	}
}
