package _435_non_overlapping_intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 组成无重叠区间的最大值
	var count = 1
	// 当前集合的最右端
	var end = intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		// 第i个区间可以加入到集合中
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}

func eraseOverlapIntervals1(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var dp = make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return n - max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
