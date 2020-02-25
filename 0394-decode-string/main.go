package _394_decode_string

import (
	`strings`
)

/*
	字符串解码: 数字[字符串]=> 字符串*数组

	因为数字和[]可以穿插出现，所以借用栈保持当前最小解码子串
	用一个栈保存数字，一个栈保存“[”来之前的的字符串，每当遇到“]”，将弹栈得到重复次数和子串，将子串解码结果保存到res中
*/
func decodeString(s string) string {
	var stackNum []int
	var stackStr []string
	var res string
	var num int
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = num*10 + int(s[i]-'0')
		case s[i] == '[':
			stackNum = append(stackNum, num)
			num = 0
			stackStr = append(stackStr, res)
			res = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]

			repeat := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]

			res = tmp + strings.Repeat(res, repeat)
		default:
			res += string(s[i])
		}
	}
	return res
}
