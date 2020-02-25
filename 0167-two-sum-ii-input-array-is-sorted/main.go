package _167_two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	var ans []int
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			ans = append(ans, i+1, j+1)
			break
		}
	}
	return ans
}
