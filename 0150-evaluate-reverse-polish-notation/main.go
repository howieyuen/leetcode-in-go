package _150_evaluate_reverse_polish_notation

import (
	`strconv`
)

func evalRPN(tokens []string) int {
	var nums []int
	index := -1
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			num1,num2 := nums[index-1],nums[index]
			nums = nums[:index-1]
			index -= 2
			if token == "+" {
				num1 += num2
			} else if token == "-" {
				num1 -= num2
			} else if token == "*" {
				num1 *= num2
			} else {
				num1 /= num2
			}
			nums = append(nums, num1)
			index++
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
			index++
		}
	}
	return nums[0]
}
