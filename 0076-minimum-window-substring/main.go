package _076_minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// hash里v存的是t中各个元素的频数
	// 右窗口每次经过一个元素都要对hash中的v-1
	// 当t中的某个元素在窗口中全部出现，这时该元素对应的v=0
	// 再多出一个相同元素时就会和不存在t中的元素一样，其v=-1
	// 这时左窗口就会移动，再逐个将变为-1的v加1，表示其不在窗口内，
	// 直到左窗口移动到与右窗口重叠或窗口中未完全包含的t中元素时，右窗口继续移动
	hash := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		for l <= r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}
	return results
}

func minWindow1(s string, t string) string {
	origin, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		origin[t[i]]++
	}
	
	n := len(s)
	length := math.MaxInt32
	left, right := -1, -1
	
	check := func() bool {
		for k, v := range origin {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for i, j := 0, 0; j < n; j++ {
		if j < n && origin[s[j]] > 0 {
			cnt[s[j]]++
		}
		for check() && i <= j {
			if j-i+1 < length {
				length = j - i + 1
				left, right = i, i+length
			}
			if _, ok := origin[s[i]]; ok {
				cnt[s[i]] -= 1
			}
			i++
		}
	}
	if left == -1 {
		return ""
	}
	return s[left:right]
}
