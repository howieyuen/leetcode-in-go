package _067_add_binary

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addBinary(a string, b string) string {
	if a == "" {
		return b
	} else if b == "" {
		return a
	}
	var ans []byte
	var carry byte = 0
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		tmp := a[i] - '0' + b[j] - '0' + carry
		carry = tmp / 2
		tmp = tmp % 2
		ans = append(ans, tmp+'0')
	}
	for ; i >= 0; i-- {
		tmp := a[i] - '0' + carry
		carry = tmp / 2
		tmp = tmp % 2
		ans = append(ans, tmp+'0')
	}
	for ; j >= 0; j-- {
		tmp := b[j] - '0' + carry
		carry = tmp / 2
		tmp = tmp % 2
		ans = append(ans, tmp+'0')
	}
	if carry == 1 {
		ans = append(ans, '1')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
