package _363_max_sum_of_rectangle_no_larger_than_k

import (
	"math"
	"sort"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	var ans = math.MinInt32
	for l := 0; l < cols; l++ {
		var rowSum = make([]int, rows)
		for r := l; r < cols; r++ {
			for i := 0; i < rows; i++ {
				rowSum[i] += matrix[i][r]
			}
			ans = max(ans, cellingK(rowSum, k))
			if ans == k {
				return ans
			}
		}
	}
	return ans
}

func cellingK(arr []int, k int) int {
	rollSum, rollMax := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if rollSum > 0 {
			rollSum += arr[i]
		} else {
			rollSum = arr[i]
		}
		rollMax = max(rollMax, rollSum)
	}
	if rollMax <= k {
		return rollMax
	}
	var ret = math.MinInt32
	for l := 0; l < len(arr); l++ {
		sum := 0
		for r := l; r < len(arr); r++ {
			sum += arr[r]
			if sum > ret && sum <= k {
				ret = sum
			}
			if ret == k {
				return ret
			}
		}
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSumSubmatrix1(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	var preSum = make([][]int, rows+1)
	for i := range preSum {
		preSum[i] = make([]int, cols+1)
	}
	// 前缀和
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	var ans = math.MinInt32
	for y1 := 0; y1 < cols; y1++ {
		for y2 := y1 + 1; y2 <= cols; y2++ {
			var set []int
			var exists = make(map[int]bool)
			set = append(set, 0)
			exists[0] = true
			for x2 := 1; x2 <= rows; x2++ {
				sum := preSum[x2][y2] - preSum[0][y2] - preSum[x2][y1] + preSum[0][y1]
				index := sort.Search(len(set), func(i int) bool {
					return set[i] >= sum-k
				})
				if index < len(set) && set[index] >= sum-k {
					if sum-set[index] > ans {
						ans = sum - set[index]
					}
				}
				if !exists[sum] {
					index := sort.Search(len(set), func(i int) bool {
						return set[i] >= sum
					})
					tmp := make([]int, 0, len(set)+1)
					tmp = append(tmp, set[:index]...)
					tmp = append(tmp, sum)
					tmp = append(tmp, set[index:]...)
					exists[sum] = true
					set = tmp
				}
			}
		}
	}
	return ans
}
