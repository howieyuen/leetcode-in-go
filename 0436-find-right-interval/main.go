package _436_find_right_interval

import (
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	var hash = make(map[int]int, n)
	for i := range intervals {
		hash[intervals[i][0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans = make([]int, len(intervals))
	for i := range intervals {
		originIndex := hash[intervals[i][0]]
		ans[originIndex] = -1
		l, r := 0, n-1
		if intervals[r][0] < intervals[i][1] {
			continue
		}
		for l < r {
			mid := (l + r) / 2
			if intervals[mid][0] >= intervals[i][1] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		ans[originIndex] = hash[intervals[l][0]]
	}
	return ans
}
