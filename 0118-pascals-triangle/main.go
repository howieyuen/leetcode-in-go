package _118_pascals_triangle

func generate(numRows int) [][]int {
	var ans [][]int
	for i := 0; i < numRows; i++ {
		var level = make([]int, i+1)
		level[0] = 1
		level[i] = 1
		for j := 1; j < i; j++ {
			level[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, level)
	}
	return ans
}

func generate1(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
				continue
			}
			tmp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans[i] = tmp
	}
	return ans
}
