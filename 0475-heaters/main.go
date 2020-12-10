package _475_heaters

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	var ans int
	for _, house := range houses {
		var curMin = math.MaxInt32
		for _, heater := range heaters {
			curMin = min(curMin, abs(heater-house))
		}
		ans = max(ans, curMin)
	}
	return ans
}

func findRadius1(houses []int, heaters []int) int {
	sort.Ints(heaters)
	var ans int
	for _, house := range houses {
		index := sort.SearchInts(heaters, house)
		var l, r int
		if index == 0 {
			l = 0
			r = 0
		} else if index == len(heaters) {
			l = index - 1
			r = index - 1
		} else {
			l = index - 1
			r = index
		}
		ans = max(ans, min(abs(heaters[l]-house), abs(heaters[r]-house)))
	}
	return ans
}

func findRadius2(houses []int, heaters []int) int {
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		a := find(heaters, house)
		res = max(res, a)
	}
	return res
}

func find(heaters []int, house int) int {
	if len(heaters) == 1 {
		return abs(house - heaters[0])
	}
	l, r := 0, len(heaters)
	for l < r {
		mid := l + (r-l)>>1
		if heaters[mid] < house {
			l = mid + 1
		} else if heaters[mid] > house {
			r = mid
		} else {
			return 0
		}
	}

	a, b := heaters[0], heaters[len(heaters)-1]
	if l < len(heaters) {
		b = heaters[l]
	}
	if l > 0 {
		a = heaters[l-1]
	}
	return min(abs(a-house), abs(b-house))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
