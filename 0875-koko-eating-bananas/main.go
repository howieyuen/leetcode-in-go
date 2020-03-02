package _875_koko_eating_bananas

func minEatingSpeed(piles []int, H int) int {
	min := 1
	max := getMax(piles)
	for min < max {
		mid := (min + max) / 2
		if hoursInNeed(piles, mid) <= H {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return min
}

func hoursInNeed(piles []int, speed int) int {
	hours := 0
	for _, pile := range piles {
		hours += (pile + speed - 1) / speed
	}
	return hours
}

func getMax(piles []int) int {
	max := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}
	return max
}
