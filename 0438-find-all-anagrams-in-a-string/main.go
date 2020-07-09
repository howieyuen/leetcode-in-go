package _438_find_all_anagrams_in_a_string

import (
	"reflect"
)

func findAnagrams(s string, p string) []int {
	windowSize, totalSize := len(p), len(s)
	if windowSize > totalSize {
		return nil
	}
	var hash1 = make([]int, 26)
	var hash2 = make([]int, 26)
	for i := 0; i < windowSize; i++ {
		hash1[s[i]-'a']++
		hash2[p[i]-'a']++
	}
	count := 0
	for i := 0; i < 26; i++ {
		if hash1[i] == hash2[i] {
			count++
		}
	}
	var ans []int
	for i := 0; i < totalSize-windowSize; i++ {
		if count == 26 {
			ans = append(ans, i)
		}
		left, right := s[i]-'a', s[i+windowSize]-'a'
		hash1[right]++
		if hash1[right] == hash2[right] {
			count++
		} else if hash1[right] == hash2[right]+1 {
			count--
		}
		hash1[left]--
		if hash1[left] == hash2[left] {
			count++
		} else if hash1[left] == hash2[left]-1 {
			count--
		}
	}
	if count == 26 {
		ans = append(ans, totalSize-windowSize)
	}
	return ans
}

func findAnagrams1(s string, p string) []int {
	windowSize, totalSize := len(p), len(s)
	if windowSize > totalSize {
		return nil
	}
	var hash1 = make([]int, 26)
	var hash2 = make([]int, 26)
	for i := 0; i < windowSize; i++ {
		hash1[s[i]-'a']++
		hash2[p[i]-'a']++
	}
	var ans []int
	for i := 0; i < totalSize-windowSize; i++ {
		if reflect.DeepEqual(hash1, hash2) {
			ans = append(ans, i)
		}
		hash1[s[i]-'a']--
		hash1[s[i+windowSize]-'a']++
	}
	if reflect.DeepEqual(hash1, hash2) {
		ans = append(ans, totalSize-windowSize)
	}
	return ans
}
