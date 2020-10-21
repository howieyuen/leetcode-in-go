package _925_long_pressed_name

// 不要用name作为基准，判断typed中每个字符的合法性
// 反向思考，优先遍历typed中子字符，判断是否与name相同或者是与自身前一个字符相同
// 最后，typed遍历完，name也应该遍历结束
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}
