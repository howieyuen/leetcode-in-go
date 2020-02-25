package _744_find_smallest_letter_greater_than_target

/*
给定一个只包含小写字母的有序数组letters 和一个目标字母 target，寻找有序数组里面「比目标字母大」的最小字母。
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)
	if letters[i] > target || letters[j-1] <= target {
		return letters[i]
	}
	for i < j {
		k := i + (j-i)/2
		if letters[k] <= target {
			i = k + 1
		} else {
			j = k
		}
	}
	return letters[i%len(letters)]
}
