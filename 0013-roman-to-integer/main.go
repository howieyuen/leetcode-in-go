package main

import (
	`fmt`
)

/**
Roman to Integer

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.

*/

func romanToInt(s string) int {
	ans := 0
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := getIntFromRoman(s[i])
		if pre > cur {
			ans -= cur
		} else {
			ans += cur
		}
		pre = cur
	}
	return ans
}

func getIntFromRoman(s uint8) int {
	switch s {
	case 'I':
		return 1;
	case 'V':
		return 5;
	case 'X':
		return 10;
	case 'L':
		return 50;
	case 'C':
		return 100;
	case 'D':
		return 500;
	case 'M':
		return 1000;
	default:
		return 0
	}
}

func main() {
	strs := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, s := range strs {
		fmt.Println(romanToInt(s))
	}
}
