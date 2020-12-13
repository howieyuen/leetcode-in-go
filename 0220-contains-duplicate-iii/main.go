package _220_contains_duplicate_iii

/*
	桶排序
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := make(map[int]int)
	w := t + 1
	for i := range nums {
		key := keyFunc(nums[i], w)
		if _, ok := buckets[key]; ok {
			return true
		}
		if val, ok := buckets[key-1]; ok && abs(nums[i]-val) < w {
			return true
		}
		if val, ok := buckets[key+1]; ok && abs(nums[i]-val) < w {
			return true
		}
		buckets[key] = nums[i]
		if i >= k {
			delete(buckets, keyFunc(nums[i-k], w))
		}
	}
	return false
}

func abs(data int) int {
	if data < 0 {
		return -data
	}
	return data
}

func keyFunc(x, w int) int {
	if x < 0 {
		return (x+1)/w - 1
	} else {
		return x / w
	}
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k; j++ {
			if i+j >= len(nums) {
				break
			}
			tmp := nums[i+j] - nums[i]
			if abs(tmp) <= t {
				return true
			}
		}
	}
	return false
}
