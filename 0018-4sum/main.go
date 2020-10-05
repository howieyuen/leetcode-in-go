package _018_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	if nums == nil || len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			c := b + 1
			d := len(nums) - 1
			for c < d {
				if c > b+1 && nums[c-1] == nums[c] {
					c++
					continue
				}
				if d < len(nums)-1 && nums[d+1] == nums[d] {
					d--
					continue
				}
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					zero := []int{nums[a], nums[b], nums[c], nums[d]}
					ans = append(ans, zero)
					c++
					d--
				} else if sum < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return ans
}

func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	var dfs func(index int, sum int, cur []int)
	dfs = func(index int, sum int, cur []int) {
		if sum == target && len(cur) == 4 {
			tmp := make([]int, 4)
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		for i := index; i < n; i++ {
			// 剩余元素不够，剪枝
			if n-i < 4-len(cur) {
				return
			}
			// 去重
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			// nums升序，剩余元素用num[i+1]填充，已经超过target，剪枝
			if i < n-1 && sum+nums[i]+(3-len(cur))*nums[i+1] > target {
				return
			}
			// nums升序，剩余元素用num[-1]填充，无法到达target，剪枝
			if i < n-1 && sum+nums[i]+(3-len(cur))*nums[n-1] < target {
				continue
			}
			cur = append(cur, nums[i])
			dfs(i+1, sum+nums[i], cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, 0, nil)
	return ans
}

func fourSum2(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < n-3; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 此时最小的sum都大于target，后面就不用查询了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// nums[i]固定，最大的sum小于target，增大nums[i]继续查询
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 此时最小的sum都大于target，后面就不用查询了
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// nums[i]和nums[j]固定，最大的sum小于target，增大nums[j]继续查询
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				// 去重
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				// 去重
				if right < n-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					temp := []int{nums[i], nums[j], nums[left], nums[right]}
					ans = append(ans, temp)
					left++
					right--
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return ans
}
