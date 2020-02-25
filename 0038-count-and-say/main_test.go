package _038_count_and_say

import (
	`fmt`
	`testing`
)

func Test_countAndSay(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(countAndSay(i))
	}
}
