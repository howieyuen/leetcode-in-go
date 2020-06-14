package _300_sum_of_mutated_array_closest_to_target

import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	
	if arr[n-1]*n <= target {
		return arr[n-1]
	}
	
	avg := target / n
	if arr[0]*n >= target {
		if target%n*2 > n {
			return avg + 1
		}
		return avg
	}
	
	var sum int
	var index int
	for index = 0; index < n; index++ {
		if arr[index] > avg {
			break
		}
		sum += arr[index]
	}
	return findBestValue(arr[index:], target-sum)
}
