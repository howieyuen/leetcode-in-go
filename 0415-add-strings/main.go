package _415_add_strings

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	bytes1, bytes2 := []byte(num1), []byte(num2)
	index1, index2 := len(num1)-1, len(num2)-1
	var carry int
	for index2 >= 0 {
		tmp := carry + int(bytes1[index1]-'0') + int(bytes2[index2]-'0')
		carry = tmp / 10
		tmp = tmp % 10
		bytes1[index1] = byte(tmp + '0')
		index1--
		index2--
	}
	for index1 >= 0 {
		if carry == 0 {
			break
		}
		tmp := carry + int(bytes1[index1]-'0')
		carry = tmp / 10
		tmp = tmp % 10
		bytes1[index1] = byte(tmp + '0')
		index1--
	}
	if carry > 0 {
		return fmt.Sprintf("%d%s", carry, bytes1)
	}
	return string(bytes1)
}

func addStrings2(num1 string, num2 string) string {
	var ans []byte
	var carry int
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		calculate(&ans, num1[i], num2[j], &carry)
	}
	for ; i >= 0; i-- {
		calculate(&ans, num1[i], '0', &carry)
	}
	for ; j >= 0; j-- {
		calculate(&ans, '0', num2[j], &carry)
	}
	if carry > 0 {
		ans = append(ans, '1')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

func calculate(ans *[]byte, a, b byte, carry *int) {
	sum := int(a-'0'+b-'0') + *carry
	*carry = sum / 10
	sum %= 10
	*ans = append(*ans, byte(sum+'0'))
}
