package test3

func minSwaps(grid [][]int) int {
	n := len(grid)
	list := make([]int, n)
	for i := range grid {
		count := 0
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] != 0 {
				break
			}
			count++
		}
		list[i] = count
	}
	
	var ans int
	for i := 0; i < n-1; i++ {
		count := n - i - 1
		j := 0
		for j < n {
			if list[j] >= count {
				ans += j - i
				list = append(list[0:j], list[j+1:]...)
				list = append([]int{0}, list...)
				break
			}
			j++
		}
		if j == n {
			return -1
		}
	}
	return ans
}
