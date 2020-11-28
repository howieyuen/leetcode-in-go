package _691_stickers_to_spell_word

import (
	"math"
)

func minStickers(stickers []string, target string) int {
	m := len(target)
	var dp = make([]int, 1<<m)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for _, s := range stickers {
		for status := range dp {
			if dp[status] == -1 {
				continue
			}
			nextStatus := status
			for i := range s {
				for j := 0; j < m; j++ {
					if s[i] == target[j] && nextStatus&(1<<j) == 0 {
						nextStatus |= 1 << j
						break
					}
				}
			}
			if dp[nextStatus] == -1 || dp[status]+1 < dp[nextStatus] {
				dp[nextStatus] = dp[status] + 1
			}
		}
	}
	return dp[1<<m-1]
}

func minStickers1(stickers []string, target string) int {
	n := len(stickers)
	count := make([][]int, n)
	for i := range stickers {
		count[i] = make([]int, 26)
		for j := range stickers[i] {
			count[i][stickers[i][j]-'a']++
		}
	}
	
	dp := make(map[string]int, 0)
	dp[""] = 0
	getMin(dp, count, target)
	return dp[target]
}

func getMin(dp map[string]int, count [][]int, target string) int {
	if t, ok := dp[target]; ok {
		return t
	}
	
	tar := make([]int, 26)
	for _, c := range target {
		tar[c-'a']++
	}
	
	res := math.MaxInt32
	for _, s := range count {
		if s[target[0]-'a'] == 0 {
			continue
		}
		t := make([]byte, 0, len(target))
		for k := 0; k < 26; k++ {
			for j := tar[k] - s[k]; j > 0; j-- {
				t = append(t, byte('a'+k))
			}
		}
		
		tmp := getMin(dp, count, string(t))
		if tmp != -1 {
			res = min(res, tmp+1)
		}
	}
	
	if res == math.MaxInt32 {
		res = -1
	}
	
	dp[target] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
