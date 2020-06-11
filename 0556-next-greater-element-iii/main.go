package _556_next_greater_element_iii

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	bits := []byte(strconv.Itoa(n))
	i := len(bits) - 2
	// 从低位到高位，找到第一个递减位置
	for i >= 0 && bits[i] >= bits[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := len(bits) - 1
	// 从低位到高位，找到第一个比bits[i]大的元素，并交换
	for j >= 0 && bits[j] <= bits[i] {
		j--
	}
	bits[i], bits[j] = bits[j], bits[i]
	for left, right := i+1, len(bits)-1; left < right; left, right = left+1, right-1 {
		bits[left], bits[right] = bits[right], bits[left]
	}
	ans, _ := strconv.Atoi(string(bits))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

func nextGreaterElement1(n int) int {
	if n < 10 || n+1 > math.MaxInt32 {
		return -1
	}
	bits := []int{}
	for n > 0 {
		bits = append(bits, n%10)
		n /= 10
	}
	index1 := -1
	for i := 1; i < len(bits); i++ {
		if bits[i] < bits[i-1] {
			index1 = i
			break
		}
	}
	if index1 == -1 {
		return -1
	}
	index2 := -1
	for i := 0; i < index1; i++ {
		if bits[i] > bits[index1] {
			index2 = i
			break
		}
	}
	bits[index1], bits[index2] = bits[index2], bits[index1]
	sort.Sort(sort.Reverse(sort.IntSlice(bits[:index1])))
	n = 0
	for i := len(bits) - 1; i >= 0; i-- {
		n = n*10 + bits[i]
	}
	if n > math.MaxInt32 || n < 0 {
		return -1
	}
	return n
}
