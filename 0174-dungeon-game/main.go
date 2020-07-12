package _174_dungeon_game

// 要保证每一步的血量>0，并且走完整个路程的血量需求为最小
// 因此从左上角到右下角，并不关心中间值的正负，动态规划结构是整个路程所需的最小血量
// 所以从右下角倒推，可以保证每一步的血量>0，并且不会因为路程中加血而影响已经走过的路
func calculateMinimumHP(dungeon [][]int) int {
	rows, cols := len(dungeon), len(dungeon[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i == rows-1 && j == cols-1 {
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == rows-1 {
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == cols-1 {
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
