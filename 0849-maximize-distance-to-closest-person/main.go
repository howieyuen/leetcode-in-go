package _849_maximize_distance_to_closest_person

// 双指针
// 遍历所有座位 seats，找出每个空位左边最近的人和右边最近的人，更新当前空位到最近的人的距离。
func maxDistToClosest(seats []int) int {
	var ans int
	for i := range seats {
		if seats[i] == 0 {
			cur := 0
			for j := i - 1; j >= 0; j-- {
				if seats[j] == 1 {
					cur = i - j
					break
				}
			}
			for j := i + 1; j < len(seats); j++ {
				if seats[j] == 1 {
					if cur != 0 {
						cur = min(cur, j-i)
					} else {
						cur = j - i
					}
					break
				}
			}
			ans = max(ans, cur)
		}
	}
	return ans
}

// 计算座位到最近的人的最大距离
func maxDistToClosest1(seats []int) int {
	n := len(seats)

	left := make([]int, n)
	for i := range seats {
		left[i] = n
		if seats[i] == 1 {
			left[i] = 0
		} else if i > 0 {
			left[i] = left[i-1] + 1
		}
	}

	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		right[i] = n
		if seats[i] == 1 {
			right[i] = 0
		} else if i < n-1 {
			right[i] = right[i+1] + 1
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, min(left[i], right[i]))
	}
	return ans
}

// 按零分组
// 如果两人之间有连续 K 个空座位，那么其中存在至少一个座位到两边最近的人的距离为 (K+1) / 2。
func maxDistToClosest2(seats []int) int {
	lastIndex := -1
	ans := 0
	for i := range seats {
		if seats[i] == 1 {
			if lastIndex >= 0 {
				ans = max(ans, (i-lastIndex)/2)
			} else {
				ans = i
			}
			lastIndex = i
		}
	}
	ans = max(ans, len(seats)-1-lastIndex)
	return ans
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
