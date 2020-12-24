package main

import (
	"math"
)

func myAtoi(str string) int {
	first := true
	flag := 1
	ans := 0
	for i := range str {
		if first {
			if str[i] == ' ' {
				continue
			} else if str[i] == '+' {
				first = false
			} else if str[i] == '-' {
				flag = -1
				first = false
			} else if str[i] >= '0' && str[i] <= '9' {
				ans = ans*10 + int(str[i]-'0')
				first = false
			} else {
				return 0
			}
		} else {
			if str[i] >= '0' && str[i] <= '9' {
				if flag == 1 && (ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && str[i] > '7')) {
					return math.MaxInt32
				}
				if flag == -1 && (ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && str[i] > '8')) {
					return math.MinInt32
				}
				ans = ans*10 + int(str[i]-'0')
			} else {
				return flag * ans
			}
		}
	}
	return flag * ans
}
