package _1

/*
see https://leetcode-cn.com/problems/validate-stack-sequences/
*/
func validateStackSequences(pushed []int, popped []int) bool {
	size := 0
	for i, j := 0, 0; i < len(pushed); i++ {
		pushed[size] = pushed[i]
		size++
		for size != 0 && pushed[size-1] == popped[j] {
			size--
			j++
		}
	}
	return size == 0
}
