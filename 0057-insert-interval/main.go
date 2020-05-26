package _057_insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	ans = append(ans, intervals[:i]...)
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	ans = append(ans, newInterval)
	ans = append(ans, intervals[i:]...)
	return ans
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
