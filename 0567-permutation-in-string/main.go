package _567_permutation_in_string

import (
	"reflect"
)

// 优化
func checkInclusion(s1 string, s2 string) bool {
	windowSize, totalSize := len(s1), len(s2)
	if windowSize > totalSize {
		return false
	}
	var hash1 = make([]int, 26)
	var hash2 = make([]int, 26)
	for i := 0; i < windowSize; i++ {
		hash1[s1[i]-'a']++
		hash2[s2[i]-'a']++
	}
	var count int
	for i := 0; i < 26; i++ {
		if hash1[i] == hash2[i] {
			count++
		}
	}
	for i := windowSize; i < totalSize; i++ {
		if count == 26 {
			return true
		}
		left, right := s2[i-windowSize]-'a', s2[i]-'a'
		hash2[right]++
		if hash2[right] == hash1[right] {
			count++
		} else if hash2[right] == hash1[right]+1 {
			count--
		}
		hash2[left]--
		if hash2[left] == hash1[left] {
			count++
		} else if hash2[left] == hash1[left]-1 {
			count--
		}
	}
	return count == 26
}

// 滑动窗口
func checkInclusion1(s1 string, s2 string) bool {
	windowSize, totalSize := len(s1), len(s2)
	if windowSize > totalSize {
		return false
	}
	var hash1 = make([]int, 26)
	var hash2 = make([]int, 26)
	for i := range s1 {
		hash1[s1[i]-'a']++
		hash2[s2[i]-'a']++
	}
	for i := windowSize; i < totalSize; i++ {
		if reflect.DeepEqual(hash1, hash2) {
			return true
		}
		hash2[s2[i-windowSize]-'a']--
		hash2[s2[i]-'a']++
	}
	return reflect.DeepEqual(hash1, hash2)
}
