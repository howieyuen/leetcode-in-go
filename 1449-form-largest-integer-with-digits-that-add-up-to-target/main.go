package _449_form_largest_integer_with_digits_that_add_up_to_target

import (
	"math"
	"strconv"
)

// 8ms 6.9MB
func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)
	for i := 1; i < target; i++ {
		dp[i] = math.MinInt32
	}
	// dp[i]>0，表示能找到target = i的组合数
	for i := 1; i <= target; i++ {
		for _, cost := range cost {
			if i >= cost {
				dp[i] = max(dp[i], dp[i-cost]+1)
			}
		}
	}
	if dp[target] <= 0 {
		return "0"
	}
	var result string
	for target > 0 {
		// 逆序拼接字符串
		for i := len(cost); i >= 0; i-- {
			c := cost[i]
			if target >= c && dp[target] == dp[target-c]+1 {
				result += strconv.Itoa(i)
				target -= c
				break
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 148ms 78.8MB
func largestNumber1(cost []int, target int) string {
	var costToNum = make(map[int]int)
	// 相同的cost，保留下标更大的数
	for i := len(cost) - 1; i >= 0; i-- {
		if _, ok := costToNum[cost[i]]; !ok {
			costToNum[cost[i]] = i + 1
		}
	}
	
	// dp[i]表示cost = i时，最大的数的字符串形式
	var dp = make(map[int]string)
	dp[0] = ""
	for i := 1; i <= target; i++ {
		for cost, num := range costToNum {
			if str, ok := dp[i-cost]; ok && cost <= i {
				str = dp[i-cost] + strconv.Itoa(num)
				dp[i] = maxString(dp[i], str)
			}
		}
	}
	if dp[target] == "" {
		return "0"
	}
	return dp[target]
}

func maxString(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	if len(a) < len(b) {
		return b
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return a
		} else if a[i] < b[i] {
			return b
		}
	}
	return a
}

// 84ms 17.5mb
func largestNumber2(cost []int, target int) string {
	var dp = make(map[int]string)
	dp[0] = ""
	for i := 1; i <= target; i++ {
		max := "0"
		if v, ok := dp[i]; ok {
			max = v
		}
		for j := range cost {
			if i < cost[j] {
				continue
			}
			sub := dp[i-cost[j]]
			if sub == "0" {
				continue
			}
			max = maxString(max, sub+strconv.Itoa(j+1))
		}
		dp[i] = max
	}
	return dp[target]
}

// 60ms 18.2MB
// 优化超时问题
// 1.没有对相同cost做优化
// 2.没有根据对台规划特点，对中间结果优化
func largestNumberN(cost []int, target int) string {
	var costToNum = make(map[int]int)
	// 相同的cost，保留下标更大的数
	for i := len(cost) - 1; i >= 0; i-- {
		if _, ok := costToNum[cost[i]]; !ok {
			costToNum[cost[i]] = i + 1
		}
	}
	
	var dfs func(target int, cache map[int]string) string
	dfs = func(target int, cache map[int]string) string {
		if target == 0 {
			return ""
		}
		if str, ok := cache[target]; ok {
			return str
		}
		ans := "0"
		for cost, num := range costToNum {
			if cost > target {
				continue
			}
			sub := dfs(target-cost, cache)
			if sub == "0" {
				continue
			}
			cur := sub + strconv.Itoa(num)
			ans = maxString(ans, cur)
		}
		cache[target] = ans
		return ans
	}
	return dfs(target, make(map[int]string))
}
