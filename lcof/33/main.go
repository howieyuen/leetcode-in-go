package _3

import (
	"math"
)

func verifyPostorder(postOrder []int) bool {
	var stack []int
	var pre = math.MaxInt64
	for i := len(postOrder) - 1; i >= 0; i-- {
		if postOrder[i] > pre {
			return false
		}
		for len(stack) > 0 && postOrder[i] < stack[len(stack)-1] {
			pre = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postOrder[i])
	}
	return true
}

func verifyPostorder1(postOrder []int) bool {
	if len(postOrder) < 2 {
		return true
	}
	return verify(postOrder, 0, len(postOrder)-1)
}

func verify(postOrder []int, left, right int) bool {
	if left >= right {
		return true
	}
	rootValue := postOrder[right]
	index := left
	for index < right && postOrder[index] < rootValue {
		index++
	}
	for i := index; i < right; i++ {
		if rootValue > postOrder[i] {
			return false
		}
	}
	return verify(postOrder, left, index-1) && verify(postOrder, index, right-1)
}
