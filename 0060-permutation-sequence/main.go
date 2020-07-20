package _060_permutation_sequence

import (
	"bytes"
)

/*
对与N的排列来讲，如果确定了第一个字母，后续的排列方法就是（N-1）的阶乘个
依照这样的规律，就可以确定了K的具体排列。
eg： 假如 N=3 k = 5; 为了确定首个字母是多少，如果按照上面的理论
第一个字母应该是 int result = k/(N-1)! = 5/(2*1)=2; 也就是第三个字符 3; 同理可以计算接下来的数字为 1,2;
*/
func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]byte, n)
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
		tokens[i-1] = byte(i + '0')
	}
	k--
	var b bytes.Buffer
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		b.WriteByte(tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}

func getPermutation2(n int, k int) string {
	var nums = make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	
	var serial = make([]int, 0, n)
	var used = make([]bool, n)
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			return
		}
		num := factorial[n-index-1]
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if num < k {
				k -= num
				continue
			}
			serial = append(serial, nums[i])
			used[i] = true
			dfs(index + 1)
		}
	}
	
	dfs(0)
	var b bytes.Buffer
	for i := range serial {
		b.WriteByte(byte(serial[i] + '0'))
	}
	return b.String()
}

func getPermutation1(n int, k int) string {
	var input = make([]byte, n)
	for i := 0; i < n; i++ {
		input[i] = byte(i + 1 + '0')
	}
	for k > 1 {
		nextPermutation(input)
		k--
	}
	return string(input)
}

func nextPermutation(input []byte) {
	n := len(input)
	
	i := n - 2
	// 从后向前相邻比较，找到第一个前者比后者小，这是将要替换的数less
	for i >= 0 && input[i+1] <= input[i] {
		i--
	}
	// 如果满足此特性，即当前排列不是降序
	if i >= 0 {
		j := n - 1
		// 从后向前，找到第一个比less大的数
		for j >= 0 && input[j] <= input[i] {
			j--
		}
		// 交换位置
		input[i], input[j] = input[j], input[i]
	}
	// 把i前的数字降序排列，得到一个比给定序列稍大一点的解
	reverse(input, i+1)
}

func reverse(bytes []byte, start int) {
	for i, j := start, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

/*
全排列
*/
func getAllPermutations(n int) []string {
	var ans []string
	var start = make([]byte, n)
	for i := 0; i < n; i++ {
		start[i] = byte(i + 1 + '0')
	}
	var backtrace func(index int, n int, cur []byte)
	backtrace = func(index int, n int, cur []byte) {
		if index == n {
			ans = append(ans, string(cur))
			return
		}
		for i := index; i < n; i++ {
			cur[i], cur[index] = cur[index], cur[i]
			backtrace(index+1, n, cur)
			cur[i], cur[index] = cur[index], cur[i]
		}
	}
	backtrace(0, n, start)
	return ans
}
