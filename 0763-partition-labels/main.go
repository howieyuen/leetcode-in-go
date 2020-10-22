package _763_partition_labels

func partitionLabels(S string) []int {
	var lastIndex = make(map[int32]int)
	for i, c := range S {
		lastIndex[c] = i
	}
	
	var start, end int
	var ans []int
	for i, c := range S {
		if lastIndex[c] > end {
			end = lastIndex[c]
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
