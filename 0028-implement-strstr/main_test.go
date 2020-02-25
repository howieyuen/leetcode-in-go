package _028_implement_strstr

import (
	`fmt`
	`testing`
)

func Test_strStr(t *testing.T) {
	haystack := "hell"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
	haystack = "aaaaa"
	needle = "bba"
	fmt.Println(strStr(haystack, needle))
	haystack = "aaaaa"
	needle = ""
	fmt.Println(strStr(haystack, needle))
	haystack = ""
	needle = ""
	fmt.Println(strStr(haystack, needle))
}
