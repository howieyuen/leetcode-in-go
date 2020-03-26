package main

import (
	`fmt`
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
	if len(strs) == 0 {return ""}
	if len(strs) == 1 {return strs[0]}
	mx := 0
	for {
		for i := 1; i < len(strs); i++ {
			if mx >= len(strs[i-1]) || mx >= len(strs[i]) ||
				strs[i-1][mx] != strs[i][mx] {
				return strs[0][:mx]
			}
		}
		mx++
	}
}

func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
