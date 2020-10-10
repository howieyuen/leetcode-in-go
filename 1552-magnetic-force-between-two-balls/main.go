package _552_magnetic_force_between_two_balls

import (
	"sort"
)

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	// 检查是否存在m-1个间距，即可放置m个球
	check := func(distance int) bool {
		i := 0
		cnt := 1
		for j := 1; j < n; j++ {
			if position[j]-position[i] >= distance {
				i = j
				cnt++
			}
			if cnt >= m {
				return true
			}
		}
		return false
	}
	
	left, right := 1, (position[n-1]-position[0])/(m-1)
	ans := 1
	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return ans
}
