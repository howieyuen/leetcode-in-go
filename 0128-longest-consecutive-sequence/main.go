package _128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	var numSet = make(map[int]bool)
	for i := range nums {
		numSet[nums[i]] = true
	}
	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			tmp := num
			l := 1
			for numSet[tmp+1] {
				tmp++
				l++
			}
			if l > longest {
				longest = l
			}
		}
	}
	return longest
}
