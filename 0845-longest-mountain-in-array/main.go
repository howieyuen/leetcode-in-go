package _845_longest_mountain_in_array

func longestMountain(a []int) int {
	var ans int
	n := len(a)
	i := 0
	for i < n {
		up, down := 0, 0
		for i < n-1 && a[i] < a[i+1] {
			i++
			up++
		}
		if up > 0 {
			for i < n-1 && a[i] > a[i+1] {
				i++
				down++
			}
			if down > 0 && ans < down+up+1 {
				ans = down + up + 1
			}
		} else {
			i++
		}
	}
	return ans
}
