package test3

import (
	"sort"
	"strings"
)

func maxNumOfSubstrings(s string) []string {
	var subs = make([][]int, 26)
	for i := range subs {
		subs[i] = make([]int, 3)
		subs[i][2] = len(s) + 1
		c := byte('a' + i)
		first := strings.IndexByte(s, c)
		if first == -1 {
			continue
		}
		subs[i][0] = first
		subs[i][1] = strings.LastIndexByte(s, c)
		subs[i][2] = subs[i][1] - subs[i][0] + 1
	}
	
	var getFullSub = func(left, right int) []int {
		for j := left + 1; j <= right; j++ {
			n := s[j] - 'a'
			// 包裹的字符左右边界也被包含
			if left <= subs[n][0] && subs[n][1] <= right {
				continue
			}
			// 反之，更新
			left = min(left, subs[n][0])
			right = max(right, subs[n][1])
			j = left
		}
		return []int{left, right, right - left + 1}
	}
	
	for i := range subs {
		if subs[i][2] == len(s)+1 {
			continue
		}
		subs[i] = getFullSub(subs[i][0], subs[i][1])
	}
	
	sort.Slice(subs, func(i, j int) bool {
		return subs[i][2] < subs[j][2]
	})
	
	var ans []string
	visited := make([]bool, len(s))
	for _, sub := range subs {
		if sub[2] == len(s)+1 {
			break
		}
		check := true
		for j := sub[0]; j <= sub[1] && check; j++ {
			check = !visited[j]
		}
		if !check {
			continue
		}
		for j := sub[0]; j <= sub[1]; j++ {
			visited[j] = true
		}
		ans = append(ans, s[sub[0]:sub[1]+1])
	}
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
