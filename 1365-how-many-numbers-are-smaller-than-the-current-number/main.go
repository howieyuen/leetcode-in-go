package _365_how_many_numbers_are_smaller_than_the_current_number

func smallerNumbersThanCurrent(nums []int) []int {
	var buckets = make([]int, 101)
	for _, num := range nums {
		buckets[num]++
	}
	for i := 0; i < 100; i++ {
		buckets[i+1] += buckets[i]
	}
	var ans = make([]int, len(nums))
	for i, num := range nums {
		if num > 0 {
			ans[i] = buckets[num-1]
		}
	}
	return ans
}

func smallerNumbersThanCurrent1(nums []int) []int {
	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		var c int
		for j := 0; j < len(nums); j++ {
			if nums[j] < t {
				c++
			}
		}
		r[i] = c
	}
	return r
}