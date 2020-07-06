package _063_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = 1
	for i := 1; i < rows; i++ {
		if obstacleGrid[i-1][0] == 1 && obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for j := 1; j < cols; j++ {
		if obstacleGrid[0][j-1] == 1 && obstacleGrid[0][j] == 0 {
			obstacleGrid[0][j] = 1
		} else {
			obstacleGrid[0][j] = 0
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[rows-1][cols-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
