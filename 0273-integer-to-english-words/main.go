package _273_integer_to_english_words

import (
	"fmt"
)

/*
将非负整数转换为其对应的英文表示

输入: 123
输出: "One Hundred Twenty Three"

输入: 12345
输出: "Twelve Thousand Three Hundred Forty Five"

输入: 1234567
输出: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

输入: 1234567891
输出: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
*/
func numberToWords(num int) string {
	if num == 0 {
		return single[num]
	}
	billions := num / BillionBase
	ans := ""
	if billions != 0 {
		ans = fmt.Sprintf("%s %s", overThousand(billions), BillionStr)
	}
	millions := num % BillionBase / MillionBase
	if millions != 0 {
		if ans != "" {
			ans = fmt.Sprintf("%s %s %s", ans, overThousand(millions), MillionStr)
		} else {
			ans = fmt.Sprintf("%s %s", overThousand(millions), MillionStr)
		}
	}
	thousands := num % MillionBase / ThousandBase
	if thousands != 0 {
		if ans != "" {
			ans = fmt.Sprintf("%s %s %s", ans, overThousand(thousands), ThousandStr)
		} else {
			ans = fmt.Sprintf("%s %s", overThousand(thousands), ThousandStr)
		}
	}
	left := num % ThousandBase
	if left != 0 {
		if ans != "" {
			ans = fmt.Sprintf("%s %s", ans, overThousand(left))
		} else {
			ans = fmt.Sprintf("%s", overThousand(left))
		}
	}
	return ans
}

func overThousand(num int) string {
	if num < ThousandBase {
		return overHundred(num)
	}
	thousands := num / ThousandBase
	left := num % ThousandBase
	if left == 0 {
		return fmt.Sprintf("%s %s", single[thousands], ThousandStr)
	} else {
		return fmt.Sprintf("%s %s %s", single[thousands], ThousandStr, overHundred(left))
	}
}

func overHundred(num int) string {
	if num < HundredBase {
		return over20(num)
	}
	hundreds := num / HundredBase
	leave := num % HundredBase
	if leave == 0 {
		return fmt.Sprintf("%s %s", single[hundreds], HundredStr)
	} else {
		return fmt.Sprintf("%s %s %s", single[hundreds], HundredStr, over20(leave))
	}
}

func over20(num int) string {
	if num < 10 {
		return single[num]
	} else if num < 20 {
		return less20[num]
	}
	tmp := num % 10
	if tmp != 0 {
		return dozens[num-tmp] + " " + single[tmp]
	} else {
		return dozens[num]
	}
}

var (
	HundredStr  = "Hundred"
	ThousandStr = "Thousand"
	MillionStr  = "Million"
	BillionStr  = "Billion"
)

var (
	HundredBase  = 100
	ThousandBase = 1000
	MillionBase  = 1000000
	BillionBase  = 1000000000
)

var dozens = map[int]string{
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

var less20 = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var single = map[int]string{
	0: "Zero",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}
