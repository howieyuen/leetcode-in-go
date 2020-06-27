package _111_maximum_nesting_depth_of_two_valid_parentheses_strings

// 你需要从中选出任意一组有效括号字符串 A 和 B，使 max(depth(A), depth(B)) 的可能取值最小
// 每个左括号都对应一个深度，而这个左括号，要么是A的，要么是B的。所以，栈上的左括号只要按奇偶分配给A和B
func maxDepthAfterSplit(seq string) []int {
	var depth = make([]int, len(seq))
	index := 0
	for i := range seq {
		if seq[i] == '(' {
			depth[index] = index & 1
		} else {
			depth[index] = (index + 1) & 1
		}
		index++
	}
	return depth
}
