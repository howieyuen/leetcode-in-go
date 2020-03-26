package main

func intToRoman(num int) string {
	var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	index := 0
	ans := ""
	for index < 13 {
		for num >= nums[index] {
			num -= nums[index]
			ans += romans[index]
		}
		index++
	}
	return ans
}
