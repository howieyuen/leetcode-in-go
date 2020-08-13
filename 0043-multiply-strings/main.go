package _043_multiply_strings

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}

func multiply1(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" || num2 == "" || num2 == "0" {
		return "0"
	}
	len1 := len(num1)
	len2 := len(num2)
	ints := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		for j := len2 - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			mul := a * b
			ints[i+j+1] += mul % 10
			ints[i+j] += mul / 10
		}
	}
	for i := len(ints) - 1; i >= 0; i-- {
		v := ints[i]
		ints[i] = v % 10
		v /= 10
		j := i - 1
		for v != 0 {
			ints[j] += v % 10
			v /= 10
			j--
		}
	}
	start := 0
	var bytes []byte
	copyValue := false
	for start < len(ints) {
		if ints[start] != 0 {
			copyValue = true
		}
		if copyValue {
			bytes = append(bytes, byte(ints[start]+'0'))
		}
		start++
	}
	return string(bytes)
}
