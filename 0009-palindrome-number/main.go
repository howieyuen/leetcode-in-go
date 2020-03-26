/**
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

package main

import (
	`fmt`
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	remain, y, tmp := 0, 0, x
	for tmp != 0 {
		remain = tmp % 10
		y = y*10 + remain
		tmp /= 10
	}
	return x == y
}

func main() {
	var x int
	for {
		if _, err := fmt.Scan(&x); err != nil {
			return
		}
		fmt.Println(isPalindrome(x))
	}
}
