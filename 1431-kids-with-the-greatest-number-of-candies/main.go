package _431_kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max = 0
	for i := range candies {
		if max < candies[i] {
			max = candies[i]
		}
	}
	var ans = make([]bool, len(candies))
	for i := range candies {
		ans[i] = candies[i]+extraCandies >= max
	}
	return ans
}
