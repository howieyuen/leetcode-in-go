package _135_candy

func candy(ratings []int) int {
	var candy = make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candy[i] < candy[i+1]+1 {
			candy[i] = candy[i+1] + 1
		}
	}
	sum := len(ratings)
	for i := range candy {
		sum += candy[i]
	}
	return sum
}
