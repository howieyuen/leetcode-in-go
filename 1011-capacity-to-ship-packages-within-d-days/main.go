package _011_capacity_to_ship_packages_within_d_days

func shipWithinDays(weights []int, D int) int {
	maxLoad := 0
	minLoad := weights[0]
	for _, weight := range weights {
		maxLoad += weight
		if minLoad < weight {
			minLoad = weight
		}
	}
	for minLoad <= maxLoad {
		tryLoad := (minLoad + maxLoad) / 2
		if days := check(weights, tryLoad); days <= D {
			maxLoad = tryLoad - 1
		} else if days > D {
			minLoad = tryLoad + 1
		}
	}
	return minLoad
}

func check(weights []int, tryLoad int) int {
	days := 0
	curLoad := 0
	for i := range weights {
		if curLoad+weights[i] <= tryLoad {
			curLoad += weights[i]
		} else {
			days++
			curLoad = weights[i]
		}
	}
	days++
	return days
}
