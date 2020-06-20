package _162_as_far_from_land_as_possible

// 每个海洋区域 (x, y)，离它最近的陆地区域到它的路径要么从上方或者左方来，要么从右方或者下方来。
// 两次动态规划，第一次从左上到右下，第二次从右下到左上
// 记 f(x, y)f(x,y) 为 (x, y)(x,y) 距离最近的陆地区域的曼哈顿距离，则我们可以推出这样的转移方程
// 第一阶段: 左上到右下
// grid[x][y]==1 f(x,y) = 0
// grid[x][y]==0 f(x,y) = min{f(x−1,y),f(x,y−1)}+1
// 第二阶段: 右下到左上
// grid[x][y]==1 f(x,y) = 0
// grid[x][y]==0 f(x,y) = min{f(x+1,y),f(x,y+1)}+1
func maxDistance(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	hasLand := false
	// 左上到右下
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				hasLand = true
				continue
			}
			if grid[i][j] == 0 {
				grid[i][j] = m + n
			}
			if i-1 >= 0 {
				grid[i][j] = min(grid[i][j], grid[i-1][j]+1) // 向下
			}
			if j-1 >= 0 {
				grid[i][j] = min(grid[i][j], grid[i][j-1]+1) // 向右
			}
			
		}
	}
	distance := 0
	// 右下到左上
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if i+1 < n {
				grid[i][j] = min(grid[i][j], grid[i+1][j]+1) // 向上
			}
			if j+1 < n {
				grid[i][j] = min(grid[i][j], grid[i][j+1]+1) // 向左
			}
			distance = max(distance, grid[i][j])
		}
	}
	if !hasLand {
		return -1
	}
	return distance - 1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 考虑最朴素的方法：
// 即求出每一个「海洋区域」(grid[i][j] == 0)的「最近陆地区域」(grid[i][j] == 1)的距离
// 最后取距离的最大值
func maxDistance1(grid [][]int) int {
	var queue [][]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	hasOcean := false
	var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var point []int
	for len(queue) > 0 {
		point = queue[0]
		x1, y1 := point[0], point[1]
		queue = queue[1:]
		for _, dir := range dirs {
			x2, y2 := x1+dir[0], y1+dir[1]
			if x2 < 0 || x2 >= len(grid) || y2 < 0 || y2 >= len(grid[x2]) || grid[x2][y2] != 0 {
				continue
			}
			grid[x2][y2] = grid[x1][y1] + 1
			hasOcean = true
			queue = append(queue, []int{x2, y2})
		}
	}
	if !hasOcean || len(point) == 0 {
		return -1
	}
	return grid[point[0]][point[1]] - 1
}
