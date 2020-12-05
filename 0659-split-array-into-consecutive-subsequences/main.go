package _659_split_array_into_consecutive_subsequences

func isPossible(nums []int) bool {
	var repeated = make(map[int]int)
	for _, num := range nums {
		repeated[num]++
	}
	var endCount = make(map[int]int)
	for _, num := range nums {
		if repeated[num] == 0 {
			continue
		}
		if endCount[num-1] > 0 {
			repeated[num]--
			endCount[num]++
			endCount[num-1]--
		} else if repeated[num+1] > 0 && repeated[num+2] > 0 {
			repeated[num]--
			repeated[num+1]--
			repeated[num+2]--
			endCount[num+2]++
		} else {
			return false
		}
	}
	return true
}
