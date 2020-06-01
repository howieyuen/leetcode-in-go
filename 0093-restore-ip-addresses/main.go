package _093_restore_ip_addresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	if len(s) < 4 {
		return ans
	}
	var dfs func(start int, cur []string)
	dfs = func(start int, cur []string) {
		if len(cur) == 4 && start < len(s) {
			return
		}
		if start == len(s) && len(cur) == 4 {
			ans = append(ans, strings.Join(cur, "."))
			return
		}
		for end := start + 1; end <= len(s); end++ {
			if num, _ := strconv.Atoi(s[start:end]); num > 255 {
				break
			}
			cur = append(cur, s[start:end])
			dfs(end, cur)
			cur = cur[:len(cur)-1]
			if s[start] == '0' {
				break
			}
		}
	}
	dfs(0, nil)
	return ans
}
