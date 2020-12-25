package _455_assign_cookies

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var index int
	for _, v := range s {
		if index < len(g) && v >= g[index] {
			index++
		}
	}
	return index
}
