package _542_01_matrix

import (
	`math`
)

/*
	0/1矩阵，找到每个元素与0的最新距离

	1) 找到所有的matrix[i][j]=0的位置，放入队列
	2) 从队列中取出，直接相连的上、下、左、右四个方向的距离，均为此位置上的距离+1
	3) 与原距离比较，有更新放回队列，无更新的从队列中删除，直到队列为空
*/

// 动态规划
func updateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	var dist = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt32
		}
	}
	// First pass: check for left and top
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
				}
				if j > 0 {
					dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
				}

			}
		}
	}

	// Second pass: check for bottom and right
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < cols-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}

		}
	}
	return dist
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 借助队列DFS
func updateMatrix1(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	rows, cols := len(matrix), len(matrix[0])
	coordinates := make([][2]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				coordinates = append(coordinates, [2]int{i, j})
				continue
			}
			matrix[i][j] = math.MaxInt32
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(coordinates) > 0 {
		curPoint := coordinates[0]
		coordinates = coordinates[1:]
		curVal := matrix[curPoint[0]][curPoint[1]]
		for _, dir := range dirs {
			i := dir[0] + curPoint[0]
			j := dir[1] + curPoint[1]
			if i < 0 || j < 0 || i == rows || j == cols || matrix[i][j] <= curVal+1 {
				continue
			}
			matrix[i][j] = curVal + 1
			coordinates = append(coordinates, [2]int{i, j})
		}
	}
	return matrix
}

// 暴力查找
func updateMatrix2(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	var dist = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
			} else {
				for a := 0; a < rows; a++ {
					for b := 0; b < cols; b++ {
						if matrix[a][b] == 0 {
							dis := minus(a, i) + minus(b, j)
							if dist[i][j] > dis {
								dist[i][j] = dis
							}
						}
					}
				}
			}
		}
	}
	return dist
}

func minus(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
