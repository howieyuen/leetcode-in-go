package _207_unique_number_of_occurrences

func uniqueOccurrences(arr []int) bool {
	times := make(map[int]int)
	for _, x := range arr {
		times[x]++
	}
	duplicates := make(map[int]bool)
	for _, val := range times {
		if duplicates[val] {
			return false
		}
		duplicates[val] = true
	}
	return true
}
