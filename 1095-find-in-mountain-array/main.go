package _095_find_in_mountain_array

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := getPeak(mountainArr)
	if target == mountainArr.get(peak) {
		return peak
	}
	if left := binarySearch(0, peak, target, mountainArr, true); left != -1 {
		return left
	}
	return binarySearch(peak+1, mountainArr.length()-1, target, mountainArr, false)
}

func binarySearch(left, right, target int, m *MountainArray, ascending bool) int {
	for left+1 < right {
		mid := left + (right-left)/2
		mountain := m.get(mid)
		if mountain == target {
			return mid
		} else if mountain < target {
			if ascending {
				left = mid
			} else {
				right = mid
			}
		} else {
			if ascending {
				right = mid
			} else {
				left = mid
			}
		}
	}
	if m.get(left) == target {
		return left
	}
	if m.get(right) == target {
		return right
	}
	return -1
}

func getPeak(m *MountainArray) int {
	left, right := 0, m.length()-1
	for left+1 < right {
		mid := left + (right-left)/2
		if m.get(mid) > m.get(mid+1) {
			right = mid
		} else {
			left = mid
		}
	}
	if m.get(left) > m.get(right) {
		return left
	}
	return right
}

type MountainArray struct {
	slice []int
}

func (m *MountainArray) get(index int) int {
	return m.slice[index]
}
func (m *MountainArray) length() int {
	return len(m.slice)
}
