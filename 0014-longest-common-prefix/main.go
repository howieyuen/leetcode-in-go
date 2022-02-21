package main

import (
	"strings"
)

/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res = strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], res) != 0 {
			res = res[0 : len(res)-1]
		}
	}
	return res
}
