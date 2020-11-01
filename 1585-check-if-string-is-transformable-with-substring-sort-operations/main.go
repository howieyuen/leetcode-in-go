package _585_check_if_string_is_transformable_with_substring_sort_operations

// s和t都是数字组成的字符串，判断能否通过排序s中的任意子字符串(升序)，将s转换成t
// 首先s和t的基数排序要一致
// 其次，要求是升序排序，则对于s中的子串，想把s[i]转换成t[j]的位置(i<=j)
// 那么比s[i]小的数都要在s[i]的右侧
// 如果满足条件，则把s[i]删除，继续判断下一个
func isTransformable(s string, t string) bool {
	// 基数队列，i表示base：0~9，j表示下标，按前后顺序排列
	var queues = make([][]int, 10)
	// 基数排序
	for i := range s {
		digit := s[i] - '0'
		queues[digit] = append(queues[digit], i)
	}
	
	for i := range t {
		digit := t[i] - '0'
		// 排序不符合
		if len(queues[digit]) == 0 {
			return false
		}
		for j := 0; j < int(digit); j++ {
			// 存在比digit小的数在s[i]的左边
			if len(queues[j]) > 0 && queues[j][0] < queues[digit][0] {
				return false
			}
		}
		queues[digit] = queues[digit][1:]
	}
	return true
}
