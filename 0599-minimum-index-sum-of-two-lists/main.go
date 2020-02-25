package _599_minimum_index_sum_of_two_lists

func findRestaurant(list1 []string, list2 []string) []string {
	var maps = make(map[string]int)
	for i := range list1 {
		maps[list1[i]] = i
	}
	var ans []string
	var curIndexSum = len(list1) + len(list2)
	for j := range list2 {
		if i, ok := maps[list2[j]]; ok {
			if j+i < curIndexSum {
				curIndexSum = j + i
				ans = []string{list2[j]}
			} else if j+i == curIndexSum {
				ans = append(ans, list2[j])
			}
		}
	}
	return ans
}
